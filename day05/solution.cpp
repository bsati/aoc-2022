#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
#include <stack>

// #define PART1

std::vector<std::string> split_string(const std::string &input, const char seperator)
{
    auto result = std::vector<std::string>();
    std::stringstream stream(input);
    std::string word;
    while (!stream.eof())
    {
        std::getline(stream, word, seperator);
        result.push_back(word);
    }
    return result;
}

void main()
{

    // read in file, line by line
    std::ifstream infile("./input.txt");
    std::string line;

    auto lines = std::vector<std::string>();

    while (std::getline(infile, line))
    {
        lines.push_back(line);
    }

    // find first empty line in input signaling state to transitions switch
    auto blank_index = -1;

    for (auto i = 0; i < lines.size(); ++i)
    {
        if (lines[i].size() == 0)
        {
            blank_index = i;
            break;
        }
    }

    auto state = std::vector<std::stack<char>>();
    const auto &labels = lines[blank_index - 1];
    for (auto i = 0; i < labels.size(); ++i)
    {
        if (labels[i] != ' ')
        {
            auto stack = std::stack<char>();
            for (auto j = blank_index - 2; j >= 0; --j)
            {
                if (lines[j][i] == ' ')
                {
                    break;
                }
                stack.push(lines[j][i]);
            }
            state.push_back(stack);
        }
    }

    for (auto i = blank_index + 1; i < lines.size(); ++i)
    {
        auto split = split_string(lines[i], ' ');
        const auto count = std::stoi(split[1]);
        auto &from = state[std::stoi(split[3]) - 1];
        auto &to = state[std::stoi(split[5]) - 1];

#ifndef PART1
        std::vector<char> elements(count);
#endif

        for (auto j = 0; j < count; ++j)
        {
#ifdef PART1
            to.push(from.top());
            from.pop();
#else
            elements[count - j - 1] = from.top();
            from.pop();
#endif
        }
#ifndef PART1
        for (auto e : elements)
        {
            to.push(e);
        }
#endif
    }

    for (auto s : state)
    {
        std::cout << s.top();
    }
}