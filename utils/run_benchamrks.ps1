$root_dir = Resolve-Path -Path ".."
echo "******** 1. standard fmt formatting lib benchmarks ******** "
go test $root_dir -bench=Fmt -benchmem -cpu 1
echo "******** 2. stringFormatter lib benchmarks ******** "
go test $root_dir -bench=Format -benchmem -cpu 1
echo "******** 3. slice fmt benchmarks ******** "
go test $root_dir -bench=SliceStandard -benchmem -cpu 1
echo "******** 4. stringFormatter lib benchmarks ******** "
go test $root_dir -bench=SliceToStringAdvanced -benchmem -cpu 1