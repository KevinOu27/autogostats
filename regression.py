import numpy as np
from scipy import stats
import time

def main():
    x = np.array([10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5])
    y = np.array([8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68])

    start_time = time.time()
    slope, intercept, r_value, p_value, std_err = stats.linregress(x, y)
    end_time = time.time()

    print(f"Linear regression results: y = {intercept:.6f} + {slope:.6f}x")
    print(f"Execution time: {end_time - start_time} seconds")

if __name__ == "__main__":
    main()
