#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <algorithm>

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

    const int width = lines.size();
    const int height = lines[0].size();

    int **heights = new int *[height];
    for (auto j = 0; j < height; ++j)
    {
        heights[j] = new int[width];
        for (auto i = 0; i < width; ++i)
        {
            heights[j][i] = lines[j][i] - '0';
        }
    }

    int directions[4][2] = {1, 0, -1, 0, 0, 1, 0, -1};
    int tree_counter = 0;
    auto scenic_scores = std::vector<int>();

    for (auto j = 0; j < height; ++j)
    {
        for (auto i = 0; i < width; ++i)
        {
            bool already_counted = false;
            int scores[4] = {0, 0, 0, 0};
            int k = 0;
            for (auto d : directions)
            {
                int dy = d[0];
                int dx = d[1];
                int y = j + dy;
                int x = i + dx;
                bool blocked = false;
                while (x >= 0 && x < width && y >= 0 && y < height)
                {
                    scores[k]++;
                    if (heights[y][x] >= heights[j][i])
                    {
                        blocked = true;
                        break;
                    }
                    x += dx;
                    y += dy;
                }
                if (!blocked && !already_counted)
                {
                    tree_counter++;
                    already_counted = true;
                }
                k++;
            }
            scenic_scores.push_back(scores[0] * scores[1] * scores[2] * scores[3]);
        }
    }

    std::cout << "part 1)" << tree_counter << "\n";
    std::sort(scenic_scores.begin(), scenic_scores.end());
    std::cout << "part 2)" << scenic_scores[scenic_scores.size() - 1] << "\n";
}