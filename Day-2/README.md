# Day 2 - Understanding Data Types and String Manipulation

## Project
Create a restaurant bill splitting program that takes in the total bill, the number of people who are splitting the bill, and the tip percentage they would like to give.  Then return each person's portion to pay.

In the original project it was specified valid tip amounts should be 10, 15, or 20%.  I would like to accept any integer as a valid percentage in this re-write.

## Description
This recreates the second day project from [100 Days of Code: The Complete Python Pro Bootcamp](https://www.udemy.com/course/100-days-of-code/) written in Go.

I am working through the class at work in Python, and on my own time am re-implmenting my solutions in Go as a way to strengthen ideas, but also learn another programming language.

## Setup & Run Instructions
![Setup Instructions and Readme Video](https://github.com/absenth/100-days-of-code/blob/main/Day-2/extras/splitbill.gif)

## Lessons Learned
Prior to working through this project I didn't realize Go won't allow you to multiply a float and int together.  As a result I changed all three variables to Float64.

Then, similar to the python program I wrote at work, I had to learn how to format the float down to 2 decimal places of precision.  This was closer to Python's syntax than I expected which was nice.

I was going to use the `math.Ceil` function to ensure we rounded up, but remembered in this case people can pay dollars and cents, unlike a previous lesson I did in Python that requires you to buy whole cans of paint, even if you only need .002 extra cans to complete the project.
