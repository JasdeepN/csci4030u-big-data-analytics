#!/bin/bash

echo "Building apriori.go"
go build apriori.go
echo "Build Complete"

arr[0]=0.75
arr[1]=0.50
arr[2]=0.35
arr[3]=0.25

for (( x = 0; x < 4; x++ )); do
	printf ${arr[x]} " ," 
	for (( i = 0; i < 5; i++ )); do
	# echo "============================================"
	# echo -e "\e[40m\e[91mBegin Run $i...\e[0m";
	# echo "Begin run..."

	start=$(date +%s.%N)

	./apriori netflix.data ${arr[x]} 0

	dur=$(echo "$(date +%s.%N) - $start" | bc)

	# echo -e "\e[1m\e[93mExecution time: \e[41m $dur \e[49m seconds\e[0m";
	# printf ">> total execution time: %.6f seconds" $dur 
	printf " %.6f " $dur 

	#forces new line
	# echo -e "\e[40m\e[91mRun Complete $i\e[0m";

done
echo ""
done

