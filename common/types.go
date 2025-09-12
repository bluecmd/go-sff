package common

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

func stringToJSON(b []byte) ([]byte, error) {
	m := map[string]interface{}{
		"value": string(b),
		"hex":   hex.EncodeToString(b),
	}
	return json.Marshal(m)
}

type String2 [2]byte
type String4 [4]byte
type String16 [16]byte

func (s String2) String() string {
	return strings.TrimSpace(string([]byte(s[:2])))
}

func (s String2) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:2]))
}

func (s *String2) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then String2 type")
	}

	*s = String2{}
	for i := 0; i < 2; i++ {
		s[i] = b[i]
	}
	return nil
}

func (s String4) String() string {
	return strings.TrimSpace(string([]byte(s[:4])))
}

func (s String4) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:4]))
}

func (s *String4) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 4 {
		return fmt.Errorf("length is shorter then String4 type")
	}

	*s = String4{}
	for i := 0; i < 4; i++ {
		s[i] = b[i]
	}
	return nil
}

func (s String16) String() string {
	return strings.TrimSpace(string([]byte(s[:16])))
}

func (s String16) MarshalJSON() ([]byte, error) {
	return stringToJSON([]byte(s[:16]))
}

func (s *String16) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 16 {
		return fmt.Errorf("length is shorter then String16 type")
	}

	*s = String16{}
	for i := 0; i < 16; i++ {
		s[i] = b[i]
	}
	return nil
}

type ValueM byte
type ValueKm byte
type Value100Mbps byte
type ValuePerc byte

func valueToJSON(b byte, unit string) ([]byte, error) {
	m := map[string]interface{}{
		"value": uint8(b),
		"unit":  unit,
		"hex":   hex.EncodeToString([]byte{byte(b)}),
	}
	return json.Marshal(m)
}

func (v ValueM) String() string {
	return fmt.Sprintf("%d m", v)
}

func (v ValueM) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "m")
}

func (v *ValueM) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = ValueM(b[0])
	return nil
}

func (v ValueKm) String() string {
	return fmt.Sprintf("%d km", v)
}

func (v ValueKm) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "km")
}

func (v *ValueKm) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = ValueKm(b[0])
	return nil
}

func (v Value100Mbps) String() string {
	return fmt.Sprintf("%d Mb/s", uint(v)*100)
}

func (v Value100Mbps) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": uint(v) * 100,
		"unit":  "Mb/s",
		"hex":   hex.EncodeToString([]byte{byte(v)}),
	}
	return json.Marshal(m)
}

func (v *Value100Mbps) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = Value100Mbps(b[0])
	return nil
}

func (v ValuePerc) String() string {
	return fmt.Sprintf("%d %%", v)
}

func (v ValuePerc) MarshalJSON() ([]byte, error) {
	return valueToJSON(byte(v), "%")
}

func (v *ValuePerc) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	*v = ValuePerc(b[0])
	return nil
}

// 16-bit unsigned integer stored in big-endian byte order
type UInt16BE [2]byte

func (u UInt16BE) Uint16() uint16 {
	return uint16(u[0])<<8 | uint16(u[1])
}

func (u UInt16BE) String() string {
	return fmt.Sprintf("%d", u.Uint16())
}

func (u UInt16BE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": u.Uint16(),
		"hex":   hex.EncodeToString([]byte{u[0], u[1]}),
	}
	return json.Marshal(m)
}

func (u *UInt16BE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then UInt16BE type")
	}

	*u = UInt16BE{b[0], b[1]}
	return nil
}

// 16-bit signed integer stored in big-endian byte order
type Int16BE [2]byte

func (i Int16BE) Int16() int16 {
	return int16(uint16(i[0])<<8 | uint16(i[1]))
}

func (i Int16BE) String() string {
	return fmt.Sprintf("%d", i.Int16())
}

func (i Int16BE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": i.Int16(),
		"hex":   hex.EncodeToString([]byte{i[0], i[1]}),
	}
	return json.Marshal(m)
}

func (i *Int16BE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then Int16BE type")
	}

	*i = Int16BE{b[0], b[1]}
	return nil
}

// Temperature in q8.8 fixed-point format stored in big-endian byte order.
// Value is a signed two's complement 16-bit integer with 8 fractional bits.
// Range: [-128.000, 127.996] °C with granularity of 1/256 °C.
type TemperatureQ8_8BE [2]byte

// Raw returns the underlying signed 16-bit value.
func (t TemperatureQ8_8BE) Raw() int16 {
	return int16(uint16(t[0])<<8 | uint16(t[1]))
}

// Celsius returns the temperature in degrees Celsius.
func (t TemperatureQ8_8BE) Celsius() float64 {
	return float64(t.Raw()) / 256.0
}

func (t TemperatureQ8_8BE) String() string {
	return fmt.Sprintf("%.3f °C", t.Celsius())
}

func (t TemperatureQ8_8BE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": t.Celsius(),
		"unit":  "°C",
		"hex":   hex.EncodeToString([]byte{t[0], t[1]}),
	}
	return json.Marshal(m)
}

func (t *TemperatureQ8_8BE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then TemperatureQ8_8BE type")
	}

	*t = TemperatureQ8_8BE{b[0], b[1]}
	return nil
}

// Optical power encoded as a 16-bit unsigned integer in big-endian byte order.
// LSB represents 0.1 µW, giving a total range of 0 to 6.5535 mW.
type PowerMilliWattBE [2]byte

// Raw returns the underlying unsigned 16-bit value.
func (p PowerMilliWattBE) Raw() uint16 {
	return uint16(p[0])<<8 | uint16(p[1])
}

// MilliWatt returns the power in milliwatt.
func (p PowerMilliWattBE) MilliWatt() float64 {
	// 1 LSB = 0.1 µW = 0.0001 mW
	return float64(p.Raw()) * 0.0001
}

// DBm returns the power in dBm.
func (p PowerMilliWattBE) DBm() float64 {
	mw := p.MilliWatt()
	return 10.0 * math.Log10(mw)
}

func (p PowerMilliWattBE) String() string {
	dBm := p.DBm()
	if math.IsInf(dBm, -1) {
		return fmt.Sprintf("%.4f mW (-inf dBm)", p.MilliWatt())
	}
	return fmt.Sprintf("%.4f mW (%.2f dBm)", p.MilliWatt(), dBm)
}

func (p PowerMilliWattBE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"mW":  p.MilliWatt(),
		"dBm": p.DBm(),
		"hex": hex.EncodeToString([]byte{p[0], p[1]}),
	}
	return json.Marshal(m)
}

func (p *PowerMilliWattBE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then PowerMilliWattBE type")
	}

	*p = PowerMilliWattBE{b[0], b[1]}
	return nil
}

// Voltage encoded as a 16-bit unsigned integer in big-endian byte order.
// LSB represents 100 µV, giving a total range of 0 to 6.5535 V.
type VoltageVoltBE [2]byte

// Raw returns the underlying unsigned 16-bit value.
func (v VoltageVoltBE) Raw() uint16 {
	return uint16(v[0])<<8 | uint16(v[1])
}

// Volts returns the voltage in volts.
func (v VoltageVoltBE) Volts() float64 {
	// 1 LSB = 100 µV = 0.0001 V
	return float64(v.Raw()) * 0.0001
}

func (v VoltageVoltBE) String() string {
	return fmt.Sprintf("%.4f V", v.Volts())
}

func (v VoltageVoltBE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"V":   v.Volts(),
		"hex": hex.EncodeToString([]byte{v[0], v[1]}),
	}
	return json.Marshal(m)
}

func (v *VoltageVoltBE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then VoltageVoltBE type")
	}

	*v = VoltageVoltBE{b[0], b[1]}
	return nil
}

// TX bias current encoded as a 16-bit unsigned integer in big-endian byte order.
// LSB represents 2 µA, giving a total range of 0 to 131.070 mA.
type CurrentMilliAmpBE [2]byte

// Raw returns the underlying unsigned 16-bit value.
func (c CurrentMilliAmpBE) Raw() uint16 {
	return uint16(c[0])<<8 | uint16(c[1])
}

// MilliAmp returns the current in milliamps.
func (c CurrentMilliAmpBE) MilliAmp() float64 {
	// 1 LSB = 2 µA = 0.002 mA
	return float64(c.Raw()) * 0.002
}

func (c CurrentMilliAmpBE) String() string {
	return fmt.Sprintf("%.3f mA", c.MilliAmp())
}

func (c CurrentMilliAmpBE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"mA":  c.MilliAmp(),
		"hex": hex.EncodeToString([]byte{c[0], c[1]}),
	}
	return json.Marshal(m)
}

func (c *CurrentMilliAmpBE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then CurrentMilliAmpBE type")
	}

	*c = CurrentMilliAmpBE{b[0], b[1]}
	return nil
}

type VendorOUI [3]byte

func (v VendorOUI) String() string {
	return fmt.Sprintf("%x:%x:%x", v[0], v[1], v[2])
}

func (v VendorOUI) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": v.String(),
		"hex":   hex.EncodeToString([]byte(v[:3])),
	}
	return json.Marshal(m)
}

func (v *VendorOUI) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 3 {
		return fmt.Errorf("length is shorter then VendorOUI type")
	}

	*v = VendorOUI{}
	for i := 0; i < 3; i++ {
		v[i] = b[i]
	}
	return nil
}

type DateCode [8]byte

func (d DateCode) String() string {
	return fmt.Sprintf("20%s-%s-%s", string(d[:2]), string(d[2:4]), string(d[4:6]))
}

func (d DateCode) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": d.String(),
		"hex":   hex.EncodeToString([]byte(d[:8])),
	}
	return json.Marshal(m)
}

func (d *DateCode) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 8 {
		return fmt.Errorf("length is shorter then DateCode type")
	}

	*d = DateCode{}
	for i := 0; i < 8; i++ {
		d[i] = b[i]
	}
	return nil
}

// Wavelength encoded as a 16-bit unsigned integer in big-endian byte order.
// The laser wavelength is equal to the 16-bit integer value divided by 20 in nm (units of 0.05 nm).
// Range: 0 to 3276.75 nm with granularity of 0.05 nm.
type WavelengthNanometerBE [2]byte

// Raw returns the underlying unsigned 16-bit value.
func (w WavelengthNanometerBE) Raw() uint16 {
	return uint16(w[0])<<8 | uint16(w[1])
}

// Nanometers returns the wavelength in nanometers.
func (w WavelengthNanometerBE) Nanometers() float64 {
	// 1 LSB = 0.05 nm, so divide by 20
	return float64(w.Raw()) / 20.0
}

func (w WavelengthNanometerBE) String() string {
	return fmt.Sprintf("%.1f nm", w.Nanometers())
}

func (w WavelengthNanometerBE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": w.Nanometers(),
		"unit":  "nm",
		"hex":   hex.EncodeToString([]byte{w[0], w[1]}),
	}
	return json.Marshal(m)
}

func (w *WavelengthNanometerBE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then WavelengthNanometerBE type")
	}

	*w = WavelengthNanometerBE{b[0], b[1]}
	return nil
}


// Tolerance encoded as a 16-bit unsigned integer in big-endian byte order.
// The tolerance is equal to the 16-bit integer value divided by 200 in nm (units of 0.005 nm).
type ToleranceNanometerBE [2]byte

// Raw returns the underlying unsigned 16-bit value.
func (w ToleranceNanometerBE) Raw() uint16 {
	return uint16(w[0])<<8 | uint16(w[1])
}

// Nanometers returns the wavelength in nanometers.
func (w ToleranceNanometerBE) Nanometers() float64 {
	// 1 LSB = 0.005 nm, so divide by 200
	return float64(w.Raw()) / 200.0
}

func (w ToleranceNanometerBE) String() string {
	return fmt.Sprintf("%.1f nm", w.Nanometers())
}

func (w ToleranceNanometerBE) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"value": w.Nanometers(),
		"unit":  "nm",
		"hex":   hex.EncodeToString([]byte{w[0], w[1]}),
	}
	return json.Marshal(m)
}

func (w *ToleranceNanometerBE) UnmarshalJSON(in []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(in, &m)
	if err != nil {
		return err
	}

	b, err := hex.DecodeString(m["hex"].(string))
	if err != nil {
		return err
	}

	if len(b) < 2 {
		return fmt.Errorf("length is shorter then ToleranceNanometerBE type")
	}

	*w = ToleranceNanometerBE{b[0], b[1]}
	return nil
}
