$root_dir = Resolve-Path -Path ".."
echo "******** 1. standadrd fmt formatting lib benchmarks ******** "
go test $root_dir -bench=Fmt -benchmem -cpu 1
echo "******** 2. stringFormatter lib benchmarks ******** "
go test $root_dir -bench=Format -benchmem -cpu 1