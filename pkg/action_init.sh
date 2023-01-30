if [[ "$PROMPT_COMMAND" == *#CMD#* ]]; then
	exit 0
fi

#TIMER_VAR#=0

#CMD#() {	
    local tcc_elapsed=$((SECONDS - #TIMER_VAR#))
    if [[ "$#TIMER_VAR#" -eq "0" || "$tcc_elapsed" -gt "#PERIOD#" ]]; then
        #TIMER_VAR#=$SECONDS		
        #THIS_PATH# #ARGS#
    fi
}

PROMPT_COMMAND="#CMD# $PROMPT_COMMAND"
#ALIAS#
echo "#MESSAGE#"
