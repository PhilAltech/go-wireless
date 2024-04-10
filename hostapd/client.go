package hostapd

import (
	"context"
	"io"
	"log"
	"time"
	"github.com/PhilAltech/go-wireless/common"
)

// HOSTPADConn is an interface to the connection
type HOSTPADConn interface {
	SendCommand(...string) (string, error)
	SendCommandBool(...string) error
	SendCommandInt(...string) (int, error)
	SendCommandWithContext(context.Context, ...string) (string, error)
	SendCommandBoolWithContext(context.Context, ...string) error
	SendCommandIntWithContext(context.Context, ...string) (int, error)
	io.Closer
	Subscribe(...string) *common.Subscription
}

// Client represents a wireless client
type Client struct {
	conn        HOSTPADConn
	ScanTimeout time.Duration
	CmdTimeout  time.Duration
	ctx         context.Context
}

// NewClient will create a new client by connecting to the
// given socket in HOSTAPD, if the socket is empty it will
// connect to the default socket /var/run/hostapd
func NewClient(iface string) (c *Client, err error) {
	c = new(Client)
	c.conn, err = common.DialHostapd(iface)
	if err != nil {
		return
	}
	c.CmdTimeout = time.Second
	return
}

// NewClientFromConn returns a new client from an already established connection
func NewClientFromConn(conn HOSTPADConn) (c *Client) {
	c = new(Client)
	c.conn = conn
	return
}

func (cl *Client) getContext() (context.Context, func()) {
	if cl.ctx != nil {
		return cl.ctx, func() {}
	}

	return context.WithTimeout(context.Background(), cl.CmdTimeout)
}

func (cl *Client) WithContext(ctx context.Context, ops ...func(wc *Client)) {
	if len(ops) == 0 {
		cl.ctx = ctx
		return
	}

	wc := NewClientFromConn(cl.conn)
	wc.ctx = ctx
	for _, op := range ops {
		op(wc)
	}
}

// Close will close the client connection
func (cl *Client) Close() {
	err := cl.conn.Close()
	if err != nil {
		log.Println("ERROR: client failed to close conn:", err)
	}
}

// Conn will return the underlying connection
func (cl *Client) Conn() *common.Conn {
	return cl.conn.(*common.Conn)
}

// Subscribe will subscribe to certain events that happen in WPA
func (cl *Client) Subscribe(topics ...string) *common.Subscription {
	return cl.conn.Subscribe(topics...)
}

// Status will return the current state of the HOSTAPD
func (cl *Client) Status() (State, error) {
	ctx, cancel := cl.getContext()
	defer cancel()
	data, err := cl.conn.SendCommandWithContext(ctx, common.CmdHostapdStatus)
	if err != nil {
		return State{}, err
	}
	s := NewState(data)
	return s, nil
}

// Enable will enable or disable the AP
func (cl *Client) Enable(enable bool) error {
	ctx, cancel := cl.getContext()
	defer cancel()
	var cmd string
	if enable {
		cmd = common.CmdHostapdEnable
	} else {
		cmd = common.CmdHostapdDisable
	}
	return cl.conn.SendCommandBoolWithContext(ctx, cmd)
}

// Reload will reload the configuration
func (cl *Client) Reload() error {
	ctx, cancel := cl.getContext()
	defer cancel()
	return cl.conn.SendCommandBoolWithContext(ctx, common.CmdHostapdReload)
}
