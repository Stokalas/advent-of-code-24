# Advent of Code '24

This is my attempt to solve at least some tasks from the challenge.

I've chosen _Go_ programming language to challenge myself even further as I have only basic knowledge of it.

## Disclaimer
This is intended for my own learning purposes, so do not take code from this repo as necessarily "good examples". 🙂

## Reflection
### Day 1
Quite fun and small tasks. It was good to remember file reading and parsing.

__Key learnings:__
- Initially, I tried reading through the entire file to count the rows before making a second pass to extract the data. This approach allowed me to allocate an array with the exact size needed, avoiding dynamic reallocations. However, it turns out that Go handles slice resizing efficiently, with the number of reallocations growing logarithmically rather than linearly. As a result, providing a reasonable guess for the initial size should be sufficient in this scenario.
- I'm not super happy with solution of the 1st task as it uses workaround of assigning _-1_ instead of used values. This works with assumption that data contains only positive integers, but would "break" if that was not the case. _Maybe I'll come back to improve it_

### Day 2
The first task was swift and fun. Though I struggled a lot with the second one. I've tried to figure out all the edge cases and then handle it with conditional statements. As I was not able to get to the result - decided to use a bit more "brute force based" solution. Felt too tired today to try to figure it out 😕 Maybe it will go better tomorrow.

__Key Learnings:__
- Usage of console args in Go.

### Day 3
It went much better today! All the time I felt in control and enjoyed implementing the solutions (no failed attempts!) Hopefully the trend stays like this tomorrow as well :). It was fun to play around with strings and their splicing. No major discoveries, though.

### Day 4
Not bad at all. Maybe a bit messy code, but had little time to solve today's tasks, so happy I did it at all... BTW Spotify Wrapped is out!

__Key Learnings:__
- str[i] for a string gives you the byte at index i, not the rune;
- comparing string elements to runes.

### Day 5
Life happens, so completing this on 7th of December. It was nice to use `Map`s and it felt quite good solving these after having a bit of rest from it.

### Day 6
Second day solved on 7th of December. I enjoyed "gamified" nature of the task. I'm not super happy with 2nd task implementation, but needed to get it done, so I can move on with other days. Brute force solution is not optimal, but at least is able to find the result.

__Key Learnings:__
- Working with arrays, them being reference type, and copying array data to another instance.

### Day 7
One week of tasks done! Today was another fun one, I've got a chance to use custom struct (though it was not necessary). I feel like I've had a few interesting ideas - using binary and trinary addition logic to generate different combinations of operands as well as concatenating numbers without using strings. Small note: this day was solved on 8th of December (still catching up! 😓)