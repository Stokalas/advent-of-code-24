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
