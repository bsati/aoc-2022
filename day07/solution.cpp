#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
#include <algorithm>

struct File
{
    std::string filename;
    int filesize;

    File(std::string f, int size) : filename(f), filesize(size) {}
};

class Directory
{
public:
    Directory *parent;
    std::vector<Directory> subdirectories;
    std::vector<File> files;
    std::string name;
    int size;

    Directory(Directory *parent, std::string name, int size) : parent(parent), name(name), size(size)
    {
        subdirectories = std::vector<Directory>();
        files = std::vector<File>();
    }

    void addFile(File file)
    {
        files.push_back(file);
        size += file.filesize;
        auto *p = parent;
        while (p != nullptr)
        {
            p->size += file.filesize;
            p = p->parent;
        }
    }
};

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

void sumDirectories(Directory &node, int *sum)
{
    if (node.size <= 100000)
    {
        *sum += node.size;
    }
    for (auto subdir : node.subdirectories)
    {
        sumDirectories(subdir, sum);
    }
}

void findDeleteCandidate(Directory &node, std::vector<int> &candidates, int neededSpace)
{
    if (node.size >= neededSpace)
    {
        candidates.push_back(node.size);
    }
    for (auto subdir : node.subdirectories)
    {
        findDeleteCandidate(subdir, candidates, neededSpace);
    }
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

    Directory root(nullptr, "/", 0);
    Directory *currentDir = &root;
    for (auto i = 1; i < lines.size(); ++i)
    {
        if (lines[i][0] == '$')
        {
            if (lines[i][2] == 'l')
            {
                i++;
                while (i < lines.size())
                {
                    if (lines[i][0] != 'd')
                    {
                        auto split = split_string(lines[i], ' ');
                        const auto size = std::stoi(split[0]);
                        const auto filename = split[1];
                        auto file = File(filename, size);
                        currentDir->addFile(file);
                    }
                    if (i + 1 < lines.size() && lines[i + 1][0] == '$')
                    {
                        break;
                    }
                    i++;
                }
            }
            else
            {
                auto dir = split_string(lines[i], ' ')[2];
                if (dir == "..")
                {
                    currentDir = currentDir->parent;
                }
                else
                {
                    currentDir->subdirectories.push_back(Directory(currentDir, dir, 0));
                    auto *newDir = &currentDir->subdirectories[currentDir->subdirectories.size() - 1];
                    currentDir = newDir;
                }
            }
        }
    }

    // Part 1
    int result = 0;
    sumDirectories(root, &result);

    std::cout << result << std::endl;

    // Part 2
    std::vector<int> deleteCandidates{};
    findDeleteCandidate(root, deleteCandidates, root.size - 40000000);
    std::sort(deleteCandidates.begin(), deleteCandidates.end());
    std::cout << deleteCandidates[0] << std::endl;
}