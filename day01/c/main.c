#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char **argv) {
  if (argc < 2) {
    fprintf(stderr, "error no file input");
    return EXIT_FAILURE;
  }

  const char *filename = argv[1];
  FILE *f = fopen(filename, "r");
  if (f == NULL) {
    fprintf(stderr, "error unable to open file %s", filename);
    return EXIT_FAILURE;
  }

  size_t answer = 0;
  char strbuf[255];

  while (fgets(strbuf, sizeof(strbuf), f) != NULL) {
    size_t str_len = strlen(strbuf) - 1, num_count = 0;
    char digit[3] = {"\0"};

    for (size_t i = 0; i < str_len; i++) {
      const char ch = strbuf[i];

      if ('0' < ch && ch <= '9') {
        if (num_count == 0) {
          digit[0] = ch;
        } else {
          digit[1] = ch;
        }
        num_count++;
      }
    }

    if (strlen(digit) == 1) {
      digit[1] = digit[0];
    }

    strbuf[strcspn(strbuf, "\n")] = 0;
    answer += atoi(digit);
  }

  printf("%zu", answer);
  fclose(f);

  return EXIT_SUCCESS;
}
