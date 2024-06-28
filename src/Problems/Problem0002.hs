module Problems.Problem0002 (solution) where

solution :: IO ()
solution = print (solve 4000000)

solve :: Int -> Int
solve n = sum [x | x <- takeWhile (< n) fibs, x `mod` 2 == 0]

fibs :: [Int]
fibs = 1 : 2 : zipWith (+) fibs (tail fibs)
