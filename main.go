package main

import "github.com/insomnia-dreams-official/magus"

// для удобства обернул все в один метод
func main()  {
	magus.RunServer()
}

// success launches:
// curl "localhost:7777/gen_tree?max_lvl=3&n=3"
// fail launches:
// curl "localhost:7777/gen_tree?max_lvl=adsf&n=3"
// curl "localhost:7777/gen_tree?max_lvl=3&n="