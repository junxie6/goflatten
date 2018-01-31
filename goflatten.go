package goflatten

import (
	"reflect"
	"strconv"
)

func Flatten(v interface{}, data map[string]interface{}, parentStr string, isParentASliceOrStruct bool) {
	t := reflect.ValueOf(v).Elem()
	typeOfT := t.Type()

	//fmt.Printf("%#v %#v============\n", t.Kind().String(), typeOfT.Name())

	keyOrig := typeOfT.Name()

	for i := 0; i < t.NumField(); i++ {
		valueField := t.Field(i)
		typeField := typeOfT.Field(i)
		//tag := typeField.Tag.Get(tagName)

		// FieldName, FieldSysType, FieldUserType, FieldValue, Tag
		//fmt.Printf("%d. %v (%v|%v): %#v, tag: %v\n", i+1, typeField.Name, valueField.Kind().String(), valueField.Type(), valueField.Interface(), tag)

		//fmt.Printf("%v.%v\n", typeOfT.Name(), typeField.Name)

		var key string

		if isParentASliceOrStruct == false {
			key = parentStr + keyOrig + "." + typeField.Name
		} else {
			key = parentStr + "." + typeField.Name
		}

		if valueField.Kind() == reflect.Struct {
			fieldPtr := valueField.Addr()
			Flatten(fieldPtr.Interface(), data, key, true)
			continue
		}

		if valueField.Kind() == reflect.Slice || valueField.Kind() == reflect.Array {
			objKey := key // User.CreditCardArr

			for ii := 0; ii < valueField.Len(); ii++ {
				key = objKey + "[" + strconv.Itoa(ii) + "]"

				if valueField.Index(ii).Kind() == reflect.Struct {
					fieldPtr := valueField.Index(ii).Addr()
					Flatten(fieldPtr.Interface(), data, key, true)
					continue
				}

				if valueField.Index(ii).Kind() == reflect.Slice || valueField.Index(ii).Kind() == reflect.Array {
					objKey2 := key
					valueField2 := valueField.Index(ii)

					for iii := 0; iii < valueField2.Len(); iii++ {
						key = objKey2 + "[" + strconv.Itoa(iii) + "]"

						if valueField2.Index(iii).Kind() == reflect.Struct {
							fieldPtr2 := valueField2.Index(iii).Addr()
							Flatten(fieldPtr2.Interface(), data, key, true)
							continue
						}

						data[key] = valueField2.Index(iii).Interface()
					}
					continue
				}

				// For string, or int ??
				data[key] = valueField.Index(ii).Interface()

				//data[key] = ""

				//fmt.Printf("HHHH: %#v\n", valueField.Index(ii).Interface())
			}

			continue
		}

		data[key] = valueField.Interface()
	}
}
