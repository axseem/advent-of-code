#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

unsigned int part1(FILE *f) {
  unsigned int sum = 0;
  char line[64];
  while (fgets(line, sizeof(line), f)) {
    char *first = line, *last = line + strlen(line);
    while (!isdigit(*first)) first++;
    while (!isdigit(*last)) last--;
    sum += (*first - '0') * 10 + (*last - '0');
  }

  return sum;
}

bool stringToDigit(char *str, unsigned int *digit) {
  const char *DIGITS[9] = {"one", "two",   "three", "four", "five",
                           "six", "seven", "eight", "nine"};

  if (isdigit(*str)) {
    *digit = *str - '0';
    return true;
  }

  for (size_t i = 0; i < sizeof(DIGITS) / sizeof(DIGITS[0]); i++) {
    if (strncmp(str, DIGITS[i], strlen(DIGITS[i])) == 0) {
      *digit = i + 1;
      return true;
    }
  }

  return false;
}

unsigned int part2(FILE *f) {
  char line[64];
  unsigned int sum = 0, first, last, i = 0;
  while (fgets(line, sizeof(line), f)) {
    for (i = 0; !stringToDigit(&line[i], &first); i++);
    for (i = strlen(line); !stringToDigit(&line[i], &last); i--);
    sum += first * 10 + last;
  }

  return sum;
}

int main(void) {
  FILE *f = fopen("input.txt", "r");
  if (f == NULL) {
    perror("fopen error");
    exit(1);
  }

  printf("--- 2023 day 01 answer ---\n");
  printf("part 1:\t%d\n", part1(f));
  rewind(f);
  printf("part 2:\t%d\n", part2(f));

  fclose(f);
  return 0;
}
