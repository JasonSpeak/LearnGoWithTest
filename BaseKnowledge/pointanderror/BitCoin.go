package pointanderror

import "fmt"

//BitCoin defines a nick name of int
type BitCoin int

//String : implement String() defined in "fmt"
func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
