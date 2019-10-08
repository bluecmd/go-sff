package sff

import (
	"os"
	"syscall"
)

const (
	i2c_SLAVE = 0x0703
)

type I2C struct {
	rc *os.File
}

func NewI2C(path string, addr uint8) (*I2C, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	if err := ioctl(f.Fd(), i2c_SLAVE, uintptr(addr)); err != nil {
		return nil, err
	}
	return &I2C{f}, nil
}

func (i2c *I2C) Write(buf []byte) (int, error) {
	return i2c.rc.Write(buf)
}

func (i2c *I2C) WriteByte(b byte) (int, error) {
	var buf [1]byte
	buf[0] = b
	return i2c.rc.Write(buf[:])
}

func (i2c *I2C) Read(p []byte) (int, error) {
	return i2c.rc.Read(p)
}

func (i2c *I2C) Close() error {
	return i2c.rc.Close()
}

func ioctl(fd, cmd, arg uintptr) (err error) {
	_, _, e1 := syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}
