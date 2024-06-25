defmodule WordLadder do
  def solve(begin_word, end_word, word_list) do
    if !Enum.member?(word_list, end_word) do
      []
    else
      bfs([[begin_word]], word_list, end_word, [], [])
    end
  end

  defp bfs(paths, word_list, end_word, visited_words, result) do
    # Get the new word(s) - is the last value in each array in 'paths'.
    new_visited_words =
      Enum.reduce(paths, visited_words, fn path, acc ->
        [List.last(path) | acc]
      end)

    # IO.puts("************************************")
    # paths |> IO.inspect(label: "paths")
    # visited_words |> IO.inspect(label: "visited_words")
    # new_visited_words |> IO.inspect(label: "new_visited_words")

    {new_paths, new_result} =
      Enum.reduce(paths, {[], result}, fn path, {paths, result} ->
        # Words that are 1 letter different and haven't been visited yet - AKA "neighbors".
        neighbors =
          word_list
          |> Enum.filter(&one_letter_difference?(List.last(path), &1))
          |> Enum.reject(&(&1 in visited_words))

        Enum.reduce(neighbors, {paths, result}, fn neighbor, {neighbor_paths, neighor_result} ->
          new_path = path ++ [neighbor]

          if neighbor == end_word do
            {neighbor_paths, [new_path | neighor_result]}
          else
            {[new_path | neighbor_paths], neighor_result}
          end
        end)
      end)

    if new_result != [] do
      new_result
    else
      bfs(new_paths, word_list, end_word, new_visited_words, new_result)
    end
  end

  defp one_letter_difference?(word1, word2) do
    word1_list = String.to_charlist(word1)
    word2_list = String.to_charlist(word2)

    # For zip functionality - see https://hexdocs.pm/elixir/1.12/Enum.html#zip/2.
    Enum.count(Enum.zip(word1_list, word2_list), fn {c1, c2} -> c1 != c2 end) == 1
  end
end
