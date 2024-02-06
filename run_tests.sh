#!/bin/sh

# Remove existing coverage file if it exists
rm -f coverage.out

# Find all directories containing a _test.go file
test_dirs=$(find . -name "*_test.go" | sed 's|/[^/]*$||' | sort -u)

for dir in $test_dirs; do
  echo "Running tests with coverage for $dir..."
  cd "$dir"
  go test -v -coverprofile=coverage.out >coverage.out 2>&1
  if [ $? -eq 0 ]; then
    echo "$dir tests passed"
  else
    echo "$dir tests failed"
    echo "Error details:"
    grep -E "\s+FAIL\s+" coverage.out | awk '{print $2}'
  fi
  cat coverage.out >>../coverage.out
  rm coverage.out
  cd ..
done

# Show combined coverage information
echo "Combined coverage information:"
go tool cover -func=coverage.out
