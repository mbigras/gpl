$pallet = %w(foo bar baz)

def main
  foo
end

def foo
  printf "%s\n", $pallet.inspect
end

main