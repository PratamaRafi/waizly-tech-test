# No.1
def MinMaxSum(arr): 
    sorted_arr = sorted(arr) 
    min_sum = sum(sorted_arr[:-1]) 
    max_sum = sum(sorted_arr[1:]) 
    
    return min_sum, max_sum

# sample data
arr = [1, 3, 5, 7, 9] 
result = MinMaxSum(arr) 
print(result[0], result[1])

# No.2
def min_max_sum(arr):
    total_sum = sum(arr)
    min_sum = float('inf')
    max_sum = float('-inf')

    for i in range(len(arr)):
        total_sum -= arr[i]
        current_sum = sum(arr[:i] + arr[i+1:])
        min_sum = min(min_sum, current_sum)
        max_sum = max(max_sum, current_sum)

    return min_sum, max_sum

arr = list(map(int, input().split()))
result = min_max_sum(arr)
print(*result)


# No.3
for i in range(1, 6):
    current_sum = sum(range(1,6))
    result =  current_sum - i
    print(result)

# No.4 
    

# No.5 
def number_proportions(arr):

    n = len(arr)
    positive_count = 0
    negative_count = 0
    zero_count = 0

    for i in range(n):
        if arr[i] > 0:
            positive_count += 1
        elif arr[i] < 0:
            negative_count += 1
        else:
            zero_count += 1

    proportion_positive = round(positive_count / n, 6)
    proportion_negative = round(negative_count / n, 6)
    proportion_zero = round(zero_count / n, 6)

    return proportion_positive, proportion_negative, proportion_zero

# Sample Data
arr = [1, 2, -3, 0, -4, 5]
positive, negative, zero = number_proportions(arr)

print(positive)
print(negative)
print(zero)

    

# No.6
# Sample Data
arr = [-4, 3, -9, 0, 4, 1]

positive_count = 0
negative_count = 0
zero_count = 0

for i in range(len(arr)):
    if arr[i] > 0:
        positive_count += 1
    elif arr[i] < 0:
        negative_count += 1
    else:
        zero_count += 1

percent_positive = positive_count / len(arr)
percent_negative = negative_count / len(arr)
percent_zero = zero_count / len(arr)

print("Persentase positif: ", percent_positive)
print("Persentase negatif: ", percent_negative)
print("Persentase zero: ", percent_zero)
    

# No.7
def AMPM_convert(s):
    # Split the untill AM PM
    part = s.split(':')
    hours = int(part[0])
    minutes = int(part[1])
    seconds = int(part[2][:-2])
    am_pm = part[2][-2:]

    # Hanlde for not 12 PM case
    if am_pm == 'PM' and hours != 12:
        hours += 12

    # Convert the time to a 24-hour format string
    return '{:02d}:{:02d}:{:02d}'.format(hours, minutes, seconds)

print(AMPM_convert('12:01:00PM')) 
print(AMPM_convert('12:01:00AM')) 
    


# No.8
from datetime import datetime

def time_convert_to_24(time_str):
    format_str = '%I:%M:%S%p'
    dt = datetime.strptime(time_str, format_str)
    time_24_format = dt.replace(hour=dt.hour % 12 + 12 if time_str[-2:] == 'AM' else dt.hour)
    return time_24_format.strftime('%H:%M:%S')

time_str = "07:05:45PM"
print(time_convert_to_24(time_str)) # Output: 19:05:45