#!/bin/sh

O1=$(ps axuww|grep ngrok_restarter|awk '{system("kill -9 " $2)}')
echo $O1

OUTPUT=$(ps axuww|grep ngrok|awk '{print $2}')
counter=0
for pid in $OUTPUT
do
  echo $pid
  kill -9 $pid
  counter=$((counter+1))
done

if [[ "$counter" -gt 0 ]]; then
   ./ngrok_restarter
   echo "./ngrok_restarter"
fi
