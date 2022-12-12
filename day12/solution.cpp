#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <queue>

struct Node
{
    int i;
    int j;
    int steps;

    Node(int i, int j, int steps) : i(i), j(j), steps(steps) {}
};

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

    const int height = lines.size();
    const int width = lines[0].size();

    char **grid = new char *[height];
    int start_i, start_j;
    for (auto j = 0; j < height; ++j)
    {
        grid[j] = new char[width];
        for (auto i = 0; i < width; ++i)
        {
            grid[j][i] = lines[j][i];
            if (grid[j][i] == 'E')
            {
                grid[j][i] = 'z';
                start_i = i;
                start_j = j;
            }
        }
    }

    int directions[4][2] = {0, 1, 0, -1, 1, 0, -1, 0};
    std::queue<Node> q;
    q.push(Node(start_i, start_j, 0));
    std::vector<std::vector<bool>> visited(height, std::vector<bool>(width, false));
    visited[start_j][start_i] = true;

    while (!q.empty())
    {
        const auto current = q.front();
        q.pop();
#ifdef PART1
        if (grid[current.j][current.i] == 'S')
#else
        if (grid[current.j][current.i] == 'a')
#endif
        {
            std::cout << "part 1) " << current.steps << std::endl;
            break;
        }
        for (auto direction : directions)
        {
            int new_i = current.i + direction[0];
            int new_j = current.j + direction[1];
            if (new_i < 0 || new_i >= width || new_j < 0 || new_j >= height || visited[new_j][new_i] || grid[current.j][current.i] - grid[new_j][new_i] > 1 && grid[new_j][new_i] != 'S')
            {
                continue;
            }
            visited[new_j][new_i] = true;
            q.push(Node(new_i, new_j, current.steps + 1));
        }
    }
}