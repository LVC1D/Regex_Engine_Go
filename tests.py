# coding: utf-8
from hstest.stage_test import StageTest
from hstest.test_case import SimpleTestCase


class RegexTest(StageTest):
    m_cases = [
        # stage 1
        ("a", "a", "true", "Two identical patterns should return True!"),
        ("a", "b", "false", "Two different patterns should not return True!"),
        ("7", "7", "true", "Two identical patterns should return True!"),
        ("6", "7", "false", "Two different patterns should not return True!"),
        (".", "a", "true", "Don't forget that '.' is a wild-card that matches any single character."),
        ("a", ".", "false", "A period in the input string is still a literal!"),
        ("", "a", "true", "An empty regex always returns True!"),
        ("", "", "true", "An empty regex always returns True!"),
        ("a", "", "false", "A non-empty regex and an empty input string always returns False!"),
        # stage 2
        ("apple", "apple", "true", "Two identical equal-length patterns should return True!"),
        (".pple", "apple", "true", "The wild-card '.' should match any single character in a string."),
        ("appl.", "apple", "true", "The wild-card '.' should match any single character in a string."),
        (".....", "apple", "true", "The wild-card '.' should match any single character in a string."),
        ("", "apple", "true", "An empty regex always returns True!"),
        ("apple", "", "false", "A non-empty regex and an empty input string always returns False!"),
        ("apple", "peach", "false", "Two different patterns should not return True!"),
        # stage 3
        ("le", "apple", "true", "If the input string contains the regex, it should return True!"),
        ("app", "apple", "true", "If the input string contains the regex, it should return True!"),
        ("a", "apple", "true", "If the input string contains the regex, it should return True!"),
        (".", "apple", "true", "Even a single wild-card character '.' can produce a match!"),
        ("apwle", "apple", "false", "Two different patterns should not return True!"),
        ("peach", "apple", "false", "Two different patterns should not return True!"),
        # stage 4
        ('^app', 'apple', "true",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!"),
        ('le$', 'apple', "true",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('^a', 'apple', "true",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!"),
        ('.$', 'apple', "true",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('apple$', 'tasty apple', "true",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('^apple', 'apple pie', "true",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!"),
        ('^apple$', 'apple', "true",
         "A regex starting with '^' and ending with '$' should match only the enclosed literals!"),
        ('^apple$', 'tasty apple', "false",
         "A regex starting with '^' and ending with '$' should match only the enclosed literals!"),
        ('^apple$', 'apple pie', "false",
         "A regex starting with '^' and ending with '$' should match only the enclosed literals!"),
        ('app$', 'apple', "false",
         "A regex ending with '$' should match the preceding pattern only at the end of the input string!"),
        ('^le', 'apple', "false",
         "A regex starting with '^' should match the following pattern only at the beginning of the input string!"),
        # stage 5
        ("colou?r", "color", "true",
         "'?' in a regex should match the preceding character exactly 0 or 1 times!"),
        ("colou?r", "colour", "true",
         "'?' in a regex should match the preceding character exactly 0 or 1 times!"),
        ("colou?r", "colouur", "false",
         "'?' in a regex should match the preceding character exactly 0 or 1 times!"),
        ("colou*r", "color", "true",
         "'*' in a regex should match the preceding character 0 or more times!"),
        ("colou*r", "colour", "true",
         "'*' in a regex should match the preceding character 0 or more times!"),
        ("colou*r", "colouur", "true",
         "'*' in a regex should match the preceding character 0 or more times!"),
        ("colou+r", "colour", "true",
         "'+' in a regex should match the preceding character 1 or more times!"),
        ("colou+r", "color", "false",
         "'+' in a regex should match the preceding character 1 or more times!"),
        (".*", "aaa", "true",
         "The repetition operators can be combined with the wild card '.'!"),
        (".+", "aaa", "true",
         "The repetition operators can be combined with the wild card '.'!"),
        (".?", "aaa", "true",
         "The repetition operators can be combined with the wild card '.'!"),
        ("no+$", "noooooooope", "false",
         "The repetition operators can be combined with other metacharacters, like '.', '^' and '$'."),
        ("^no+", "noooooooope", "true",
         "The repetition operators can be combined with other metacharacters, like '.', '^' and '$'."),
        ("^no+pe$", "noooooooope", "true",
         "The repetition operators can be combined with other metacharacters, like '.', '^' and '$'."),
        ("^n.+pe$", "noooooooope", "true",
         "The repetition operators can be combined with other metacharacters, like '.', '^' and '$'."),
        ("^n.+p$", "noooooooope", "false",
         "The repetition operators can be combined with other metacharacters, like '.', '^' and '$'."),
        ('^.*c$', 'abcabc', "true",
         "The repetition operators can be combined with other metacharacters, like '.', '^' and '$'."),
        # stage 6
        ("\\.$", "end.", "true",
         "Don't forget that '\\' is an escape symbol in Python itself, so it has to be duplicated!"),
        ("3\\+3", "3+3=6", "true",
         "Don't forget that '\\' is an escape symbol in Python itself, so it has to be duplicated!"),
        ("\\?", "Is this working?", "true",
         "Don't forget that '\\' is an escape symbol in Python itself, so it has to be duplicated!"),
        ("\\\\", "\\", "true",
         "Don't forget that '\\' is an escape symbol in Python itself, so it has to be duplicated!"),
        ("colou\\?r", "color", "false",
         "Don't forget that '\\' is an escape symbol in Python itself, so it has to be duplicated!"),
        ("colou\\?r", "colour", "false",
         "Don't forget that '\\' is an escape symbol in Python itself, so it has to be duplicated!")
    ]

    def generate(self):
        return [
            SimpleTestCase(
                stdin="{0}|{1}".format(regex, text),
                stdout=output,
                feedback=fb
            ) for regex, text, output, fb in self.m_cases
        ]


if __name__ == '__main__':
    RegexTest().run_tests()
