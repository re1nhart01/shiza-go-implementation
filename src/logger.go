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

func LogLexerError(pos int64, row int64, content string, withFatalExit bool) {
	fmt.Println(MainErrorMessage)
	fmt.Printf("[SHIZA]: %s, on position %d:%d found an error! \n", ErrorTypes["lexical"], row, pos)
	fmt.Printf("[SHIZA]: Code path: \n" + "  " + content + "\n")
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	if withFatalExit {
		os.Exit(0)
	}
}

func LogParseError(pos int64, line int64, content string, withFatalExit bool) {
	fmt.Println(MainErrorMessage)
	fmt.Printf("[SHIZA]: %s, on position %d:%d expects a name %s! \n", ErrorTypes["parse"], pos, line, content)
	if withFatalExit {
		os.Exit(0)
	}
}

func LogBundlingError(errMsg string, withFatalExit bool) {
	fmt.Println(MainErrorMessage)
	fmt.Printf("[SHIZA]: Oops, %s. Check file permissions or restart a compilation.", errMsg)
	if withFatalExit {
		os.Exit(0)
	}
}
