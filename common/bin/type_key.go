package bin

var (
	tagUint8      = byte(0x01)
	tagUint16     = byte(0x02)
	tagUint32     = byte(0x03)
	tagUint64     = byte(0x04)
	tagBytes      = byte(0x05)
	tagString     = byte(0x06)
	tagBool       = byte(0x07)
	tagHash256    = byte(0x08)
	tagSignature  = byte(0x09)
	tagAddress    = byte(0x10)
	tagPublicKey  = byte(0x11)
	tagAmount     = byte(0x12)
	tagBigInt     = byte(0x13)
	tagAddressArr = byte(0x14)
	tagSlice      = byte(0x15)
	tagAmountArr  = byte(0x16)
)
