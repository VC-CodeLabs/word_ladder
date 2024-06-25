defmodule WordLadderTest do
  use ExUnit.Case

  defp sort_results(results) do
    Enum.map(results, &Enum.sort(&1))
    |> Enum.sort()
  end

  test "basic transformation with two shortest paths" do
    assert sort_results(
             WordLadder.solve("hit", "cog", ["hot", "dot", "dog", "lot", "log", "cog"])
           ) ==
             sort_results([
               ["hit", "hot", "dot", "dog", "cog"],
               ["hit", "hot", "lot", "log", "cog"]
             ])
  end

  test "single letter change" do
    assert sort_results(
             WordLadder.solve("a", "c", [
               "a",
               "b",
               "c",
               "d",
               "e",
               "f",
               "g",
               "h",
               "i",
               "j",
               "k",
               "l",
               "m",
               "n",
               "o",
               "p",
               "q",
               "r",
               "s",
               "t",
               "u",
               "v",
               "w",
               "x",
               "y",
               "z"
             ])
           ) ==
             sort_results([["a", "c"]])
  end

  test "no possible transformation" do
    assert sort_results(WordLadder.solve("hit", "cog", ["hot", "dot", "dog", "lot", "log"])) == []
  end

  test "very basic transformation" do
    assert sort_results(WordLadder.solve("lame", "same", ["same", "came", "lame", "name"])) ==
             sort_results([["lame", "same"]])
  end

  test "transformation with repeated characters" do
    assert sort_results(
             WordLadder.solve("aaaa", "bbbb", ["aaaa", "abbb", "aaab", "aabb", "bbbb"])
           ) ==
             sort_results([["aaaa", "aaab", "aabb", "abbb", "bbbb"]])
  end

  test "multiple valid transformations" do
    assert sort_results(
             WordLadder.solve("stone", "money", [
               "stoke",
               "stony",
               "stome",
               "stomy",
               "stoey",
               "htoey",
               "htney",
               "hiney",
               "miney",
               "ttney",
               "toney",
               "itoey",
               "mtney",
               "soney",
               "store",
               "storm",
               "story",
               "monte",
               "monny",
               "monty",
               "money",
               "stane",
               "stine",
               "maney",
               "honey",
               "monde",
               "stnny",
               "mtone",
               "mtnne",
               "monne",
               "monee",
               "stnne",
               "mtnee"
             ])
           ) ==
             sort_results([
               ["stone", "mtone", "mtnne", "monne", "monee", "money"],
               ["stone", "mtone", "mtnne", "monne", "monny", "money"],
               ["stone", "mtone", "mtnne", "mtnee", "monee", "money"],
               ["stone", "mtone", "mtnne", "mtnee", "mtney", "money"],
               ["stone", "stnne", "mtnne", "monne", "monee", "money"],
               ["stone", "stnne", "mtnne", "monne", "monny", "money"],
               ["stone", "stnne", "mtnne", "mtnee", "monee", "money"],
               ["stone", "stnne", "mtnne", "mtnee", "mtney", "money"]
             ])
  end
end
