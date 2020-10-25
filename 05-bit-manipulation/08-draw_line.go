package bitmanipulation

// A monochrome screen is stored as a single array of bytes, allowing eight
// consecutive pixels to be stored in one byte. The screen has width w,
// where wis divisible by 8 (that is, no byte will be split across rows).
// The height of the screen, of course, can be derived from the length of
// the array andthewidth.
// Implement a function that draws a horizontal line from (x1, y) to (x2, y).
//
// The method signature should look something like:
//
// 		drawLine(byte[] screen, int width, int xl1 int x2, int y)
