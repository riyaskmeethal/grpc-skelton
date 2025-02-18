#!/bin/bash


# Check if any arguments are passed
if [ "$#" -eq 0 ]; then
    echo "No arguments provided."
    echo "Check your 'proto' directory and pass the file names with extention as aurguments to excute protoc command"
    exit 1
fi

for i in "$@"
do
if [[ "$i" == *.proto ]]; then
FILE_NAME="${i%.proto}"
echo $FILE_NAME

# Create the folder structure if it doesn't exist
TARGET_DIR="./proto/generated/$FILE_NAME"
mkdir -p "$TARGET_DIR"

protoc -I ./proto -I ./proto/validate -I ./proto/google/api \
   --go_out=paths=source_relative:./proto/generated/$FILE_NAME \
   --go-grpc_out=paths=source_relative:./proto/generated/$FILE_NAME \
   --grpc-gateway_out ./proto/generated/$FILE_NAME --grpc-gateway_opt paths=source_relative,grpc_api_configuration=./proto/${FILE_NAME}.yaml \
   --validate_out=paths=source_relative,lang=go:./proto/generated/$FILE_NAME \
   ./proto/${i}

# Check if protoc command failed
if [ $? -ne 0 ]; then
    echo "protoc command failed."
else
    echo "protoc command for $i succeeded."
fi

else
echo "The provided file name $i must contain a .proto extension."
fi

done


