defmodule Codelab do
  use Application

  @impl true
  def start(_type, _args) do
    start()

    children = [
      # Starts a worker by calling: Foo.Worker.start_link(arg)
      # {Foo.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Foo.Supervisor]
    Supervisor.start_link(children, opts)
  end

  def start do
    # Set values here:
    begin_word = "hit"
    end_word = "cog"
    word_list = ["hot", "dot", "dog", "lot", "log", "cog"]

    results = WordLadder.solve(begin_word, end_word, word_list)

    IO.inspect(results)

    :ok
  end
end
