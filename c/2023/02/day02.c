#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int part1(FILE *f) {
  const int MAX_R = 12, MAX_G = 13, MAX_B = 14;
  int gameIndex = 1, sum = 0;

  char line[256];
  while (fgets(line, sizeof(line), f)) {
    bool isValid = true;

    char *i = strchr(line, ':');
    while (*i != '\n' && *i != '\0') {
      i += 2;
      int value = atoi(i);
      i = strchr(i, ' ') + 1;
      switch (*i) {
        case 'r':
          if (value > MAX_R) isValid = false;
          i += 3;
          break;
        case 'g':
          if (value > MAX_G) isValid = false;
          i += 5;
          break;
        case 'b':
          if (value > MAX_B) isValid = false;
          i += 4;
          break;
      }
      if (!isValid) break;
    }

    if (isValid) sum += gameIndex;
    gameIndex++;
  }

  return sum;
}

int part2(FILE *f) {
  int sum = 0;

  char line[256];
  while (fgets(line, sizeof(line), f)) {
    int max_r = 0, max_g = 0, max_b = 0;

    char *i = strchr(line, ':');
    while (*i != '\n' && *i != '\0') {
      i += 2;
      int value = atoi(i);
      i = strchr(i, ' ') + 1;
      switch (*i) {
        case 'r':
          if (value > max_r) max_r = value;
          i += 3;
          break;
        case 'g':
          if (value > max_g) max_g = value;
          i += 5;
          break;
        case 'b':
          if (value > max_b) max_b = value;
          i += 4;
          break;
      }
    }

    sum += max_r * max_g * max_b;
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
