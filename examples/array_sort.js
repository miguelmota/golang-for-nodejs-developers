const stringArray = ['a', 'd', 'z', 'b', 'c', 'y']
const stringArraySortedAsc = stringArray.sort((a, b) => a > b ? 1 : -1)
console.log(stringArraySortedAsc)

const stringArraySortedDesc = stringArray.sort((a, b) => a > b ? -1 : 1)
console.log(stringArraySortedDesc)


const numberArray = [1, 3, 5, 9, 4, 2, 0]
const numberArraySortedAsc = numberArray.sort((a, b) => a - b)
console.log(numberArraySortedAsc)

const numberArraySortedDesc = numberArray.sort((a, b) => b - a)
console.log(numberArraySortedDesc)

const collection = [
    {
        name: "Li L",
        age: 8
    },
    {
        name: "Json C",
        age: 3
    },
    {
        name: "Zack W",
        age: 15
    },
    {
        name: "Yi M",
        age: 2
    }
]

const collectionSortedByAgeAsc = collection.sort((a, b) => a.age - b.age)
console.log(collectionSortedByAgeAsc)

const collectionSortedByAgeDesc = collection.sort((a, b) => b.age - a.age)
console.log(collectionSortedByAgeDesc)