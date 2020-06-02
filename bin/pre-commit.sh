#!/bin/sh

STAGED_GO_FILES=$(git diff --cached --name-only | grep ".go$")
if [[ "$STAGED_GO_FILES" = "" ]]; then
 exit 0
fi

PASS=true
for FILE in $STAGED_GO_FILES
do 
    goimports -l -w $FILE
    if [[ $? == 1 ]]; then 
        PASS=false
    fi
done

if ! $PASS; then 
    echo "Some files were formatted. COMMIT UNSUCCESSFUL\n"
    exit 1
else 
    echo "COMMIT SUCCESSFUL\n"
fi

exit 0
