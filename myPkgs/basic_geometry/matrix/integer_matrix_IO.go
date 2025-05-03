package matrix

import (
	"encoding/gob"
	"fmt"
	"os"
)

/**/
func (imat *IntegerMatrix2D) Save_A_File(file_name string) error {
	file, err := os.Create(fmt.Sprintf("%s.gob", file_name))
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(imat)
	return err
}

/**/
func (imat *IntegerMatrix2D) Load_A_File(file_name string) (IntegerMatrix2D, error) {
	out_matrix := make(IntegerMatrix2D, 0)
	file, err := os.Open(fmt.Sprintf("%s.gob", file_name))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&out_matrix)
	if err != nil {
		return nil, err
	}
	return out_matrix, err
}
