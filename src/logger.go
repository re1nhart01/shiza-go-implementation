package src

import (
	"fmt"
	"os"
)

func LogFSError(content string, withFatalExit bool) {
	fmt.Println(MainErrorMessage, ":")
	fmt.Printf("Caught an error after reading file: <%s>", content)
	if withFatalExit {
		os.Exit(0)
	}
}

func LogLexerError(pos int64, content string, withFatalExit bool) {
	fmt.Println(MainErrorMessage)
	fmt.Printf("[SHIZA]: %s, on position %d found an error! \n", ErrorTypes["lexical"], pos)
	fmt.Printf("[SHIZA]: Code path: \n" + "  " + content)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	if withFatalExit {
		os.Exit(0)
	}
}

func LogParseError(pos int64, content string, withFatalExit bool) {
	fmt.Println(MainErrorMessage)
	fmt.Printf("[SHIZA]: %s, on position %d expects a name %s! \n", ErrorTypes["parse"], pos, content)
	if withFatalExit {
		os.Exit(0)
	}
}
