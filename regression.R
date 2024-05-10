x = c(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5)
y = c(8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68)

start_time = Sys.time()
fit = lm(y ~ x) #lm runs linear regression model
end_time = Sys.time()

cat(sprintf("Linear regression results: y = %f + %fx\n", coef(fit)[1], coef(fit)[2]))
cat(sprintf("Execution time: %f seconds\n", as.numeric(end_time - start_time, units="secs")))


