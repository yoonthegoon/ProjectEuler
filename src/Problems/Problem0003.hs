module Problems.Problem0003 (solution) where

solution :: IO ()
solution = print (solve 600851475143)

solve :: Int -> Int
solve n = largestPrimeFactor n limit
  where
    sr = floor (sqrt (fromIntegral n))
    limit = if even sr then sr - 1 else sr

largestPrimeFactor :: Int -> Int -> Int
largestPrimeFactor n limit  -- assumes n is not prime
  | n `mod` limit == 0 && isPrime limit = limit
  | otherwise = largestPrimeFactor n (limit - 2)  -- assumes limit is odd

isPrime :: Int -> Bool
isPrime n
  | n < 2 = False
  | n == 2 = True
  | even n = False
  | otherwise = null [x | x <- [3, 5 .. limit], n `mod` x == 0]
  where
    limit = floor (sqrt (fromIntegral n))
