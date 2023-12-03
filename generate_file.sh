#!/bin/sh

year=$1
day=$2

file="year$year/day$day"

cat << EOF > ${file}.go
package year$year

import (

)

func Day${day}Part1() {
    input := utils.ReadStringFromFile($year, $day)
}

func Day${day}Part2(){

}
EOF

touch inputs/$year-$day.txt
touch inputs/$year-$day-test.txt