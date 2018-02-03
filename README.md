# try-grpc-go

protoc -I grpc/ grpc/*.proto --go_out=plugins=grpc:grpc

# try

50000ゴルーチン実行

## actionWithEachConnect（各goroutineでの実行の都度、コネクションを生成）

以下エラー発生

could not grpc.Dial: context deadline exceeded

could not CreateMessage: rpc error: code = Unavailable desc = all SubConns are in TransientFailure

## actionWithCommonConnect（事前に生成済みのコネクションを各goroutineに渡して実行）

エラーなく終了
