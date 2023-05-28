package main

////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

var (
	delims = ":-"
	reMAC  = regexp.MustCompile(`^([0-9a-fA-F]{2}[` + delims + `]){5}([0-9a-fA-F]{2})$`)
)

////////////////////////////////////////////////////////////////////////////////

// MACAddress represents a 6 byte network mac address.
type MACAddress [6]byte

// A MagicPacket is constituted of 6 bytes of 0xFF followed by 16-groups of the
// destination MAC address.
type MagicPacket struct {
	header  [6]byte
	payload [16]MACAddress
}

// New 返回一个基于mac地址字符串的魔法包。
func New(mac string) (*MagicPacket, error) {
	var packet MagicPacket
	var macAddr MACAddress

	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}

	// We only support 6 byte MAC addresses since it is much harder to use the
	// binary.Write(...) interface when the size of the MagicPacket is dynamic.
	if !reMAC.MatchString(mac) {
		return nil, fmt.Errorf("%s is not a IEEE 802 MAC-48 address", mac)
	}

	// Copy bytes from the returned HardwareAddr -> a fixed size MACAddress.
	for idx := range macAddr {
		macAddr[idx] = hwAddr[idx]
	}

	// Setup the header which is 6 repetitions of 0xFF.
	for idx := range packet.header {
		packet.header[idx] = 0xFF
	}

	// Setup the payload which is 16 repetitions of the MAC addr.
	for idx := range packet.payload {
		packet.payload[idx] = macAddr
	}

	return &packet, nil
}

// Marshal 将魔术包结构序列化为一个102字节数组。
func (mp *MagicPacket) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, mp); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ipFromInterface returns a `*net.UDPAddr` from a network interface name.
func ipFromInterface(iface string) (*net.UDPAddr, error) {
	ief, err := net.InterfaceByName(iface)
	if err != nil {
		return nil, err
	}

	addrs, err := ief.Addrs()
	if err == nil && len(addrs) <= 0 {
		err = fmt.Errorf("no address associated with interface %s", iface)
	}
	if err != nil {
		return nil, err
	}

	// Validate that one of the addrs is a valid network IP address.
	for _, addr := range addrs {
		switch ip := addr.(type) {
		case *net.IPNet:
			if !ip.IP.IsLoopback() && ip.IP.To4() != nil {
				return &net.UDPAddr{
					IP: ip.IP,
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("no address associated with interface %s", iface)
}

// wake 执行唤醒指令
func Wake(macAddr string, ip string, port string) error {
	macAddr = strings.ReplaceAll(macAddr, ":", "-")
	macAddr = strings.ReplaceAll(macAddr, "：", "-")
	bcastInterface := ""
	bcastAddr := fmt.Sprintf("%s:%s", ip, port)
	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
	if err != nil {
		return err
	}

	// 获取魔术包
	mp, err := New(macAddr)
	if err != nil {
		return err
	}

	// 魔术包转字节流
	bs, err := mp.Marshal()
	if err != nil {
		return err
	}

	var localAddr *net.UDPAddr
	if bcastInterface != "" {
		localAddr, err = ipFromInterface(bcastInterface)
		if err != nil {
			return err
		}
	}

	// 创建UDP链接
	conn, err := net.DialUDP("udp", localAddr, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// 开始发送UDP数据
	fmt.Printf("正在发送魔术包MAC: %s 目标: %s\n", macAddr, bcastAddr)
	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("已发送的魔术包 %d 字节 (应发送魔术包 102 字节)", n)
	}
	if err != nil {
		return err
	}

	fmt.Printf("魔术包发送成功MAC: %s 目标: %s\n", macAddr, bcastAddr)
	return nil
}
