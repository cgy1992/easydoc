package main

import (
	"flag"
	"fmt"
	"github.com/wuyumin/easydoc"
	"github.com/wuyumin/easydoc/utils"
	"github.com/wuyumin/easydoc/utils/version"
)

var (
	versionPtr   *bool
	helpPtr      *bool
	initPtr      *bool
	buildPtr     *bool
	emptyDistPtr *bool
	serverPtr    *bool
	portPtr      *string
	pathPtr      *string
)

func init() {
	versionPtr = flag.Bool("version", false, "Print the version number of EasyDoc")
	helpPtr = flag.Bool("help", false, "Help about EasyDoc")
	initPtr = flag.Bool("init", false, "Init the document structure")
	buildPtr = flag.Bool("build", false, "Build the document")
	emptyDistPtr = flag.Bool("emptydist", false, "Empty dist directory")
	serverPtr = flag.Bool("server", false, "Start web server")
	portPtr = flag.String("port", "", "Web port")
	pathPtr = flag.String("path", "", "Web path")
	flag.Parse()
}

func main() {
	// Print version
	fmt.Println("")
	fmt.Println("EasyDoc", version.Version)
	fmt.Println("Author:", "Yumin Wu")
	fmt.Println("Website:", "https://easydoc.089858.com")
	fmt.Println("GitHub:", "https://github.com/wuyumin/easydoc")
	fmt.Println("")

	var err error
	switch {
	case *versionPtr:

	case *helpPtr:
		fmt.Println(`EasyDoc, Easy to generate Documents.`)
		fmt.Println("")
		flag.PrintDefaults()
	case *initPtr:
		err = easydoc.GenerateInit()
		utils.CheckErr(err)
		fmt.Println("Initialization is OK.")
	case *buildPtr:
		err = easydoc.GenerateDoc(false)
		utils.CheckErr(err)
	case *emptyDistPtr:
		err = easydoc.EmptyDist()
		utils.CheckErr(err)
	case *serverPtr:
		err = easydoc.StartServer(*portPtr, *pathPtr)
		utils.CheckErr(err)
	default:
		flag.PrintDefaults()
	}
}
