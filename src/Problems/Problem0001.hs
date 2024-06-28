module Problems.Problem0001 (solution) where

solution :: IO ()
solution = print (solve 1000)

solve :: Int -> Int
solve n = sum [x | x <- [1 .. n - 1], x `mod` 3 == 0 || x `mod` 5 == 0]
