// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package model

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type UserInfo struct {
	UserId     int64  `thrift:"user_id,1" frugal:"1,default,i64" json:"user_id"`
	Name       string `thrift:"name,2" frugal:"2,default,string" json:"name"`
	CreateTime string `thrift:"create_time,3" frugal:"3,default,string" json:"create_time"`
	UpdateTime string `thrift:"update_time,4" frugal:"4,default,string" json:"update_time"`
	IsDeleted  bool   `thrift:"is_deleted,5" frugal:"5,default,bool" json:"is_deleted"`
}

func NewUserInfo() *UserInfo {
	return &UserInfo{}
}

func (p *UserInfo) InitDefault() {
	*p = UserInfo{}
}

func (p *UserInfo) GetUserId() (v int64) {
	return p.UserId
}

func (p *UserInfo) GetName() (v string) {
	return p.Name
}

func (p *UserInfo) GetCreateTime() (v string) {
	return p.CreateTime
}

func (p *UserInfo) GetUpdateTime() (v string) {
	return p.UpdateTime
}

func (p *UserInfo) GetIsDeleted() (v bool) {
	return p.IsDeleted
}
func (p *UserInfo) SetUserId(val int64) {
	p.UserId = val
}
func (p *UserInfo) SetName(val string) {
	p.Name = val
}
func (p *UserInfo) SetCreateTime(val string) {
	p.CreateTime = val
}
func (p *UserInfo) SetUpdateTime(val string) {
	p.UpdateTime = val
}
func (p *UserInfo) SetIsDeleted(val bool) {
	p.IsDeleted = val
}

var fieldIDToName_UserInfo = map[int16]string{
	1: "user_id",
	2: "name",
	3: "create_time",
	4: "update_time",
	5: "is_deleted",
}

func (p *UserInfo) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField4(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 5:
			if fieldTypeId == thrift.BOOL {
				if err = p.ReadField5(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_UserInfo[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *UserInfo) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return err
	} else {
		p.UserId = v
	}
	return nil
}

func (p *UserInfo) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Name = v
	}
	return nil
}

func (p *UserInfo) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.CreateTime = v
	}
	return nil
}

func (p *UserInfo) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.UpdateTime = v
	}
	return nil
}

func (p *UserInfo) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return err
	} else {
		p.IsDeleted = v
	}
	return nil
}

func (p *UserInfo) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("UserInfo"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}
		if err = p.writeField4(oprot); err != nil {
			fieldId = 4
			goto WriteFieldError
		}
		if err = p.writeField5(oprot); err != nil {
			fieldId = 5
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *UserInfo) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user_id", thrift.I64, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI64(p.UserId); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *UserInfo) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *UserInfo) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("create_time", thrift.STRING, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.CreateTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *UserInfo) writeField4(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("update_time", thrift.STRING, 4); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.UpdateTime); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 4 end error: ", p), err)
}

func (p *UserInfo) writeField5(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("is_deleted", thrift.BOOL, 5); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteBool(p.IsDeleted); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 5 end error: ", p), err)
}

func (p *UserInfo) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserInfo(%+v)", *p)
}

func (p *UserInfo) DeepEqual(ano *UserInfo) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.UserId) {
		return false
	}
	if !p.Field2DeepEqual(ano.Name) {
		return false
	}
	if !p.Field3DeepEqual(ano.CreateTime) {
		return false
	}
	if !p.Field4DeepEqual(ano.UpdateTime) {
		return false
	}
	if !p.Field5DeepEqual(ano.IsDeleted) {
		return false
	}
	return true
}

func (p *UserInfo) Field1DeepEqual(src int64) bool {

	if p.UserId != src {
		return false
	}
	return true
}
func (p *UserInfo) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Name, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfo) Field3DeepEqual(src string) bool {

	if strings.Compare(p.CreateTime, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfo) Field4DeepEqual(src string) bool {

	if strings.Compare(p.UpdateTime, src) != 0 {
		return false
	}
	return true
}
func (p *UserInfo) Field5DeepEqual(src bool) bool {

	if p.IsDeleted != src {
		return false
	}
	return true
}
