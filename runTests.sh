cd /Users/werner/go-workspace/src/github.com/whihn/go-sandbox
go test ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | sed ''/ok/s//$(printf "\033[32mok\033[0m")/''
