from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.alert import Alert
import time


driver = webdriver.Firefox(executable_path="geckodriver.exe")
driver.get('https://gdsports.or.kr/login/login.do')
driver.maximize_window()
driver.implicitly_wait(15)


# 로그인
driver.find_element(By.ID, 'userId').send_keys('bigster31')
driver.find_element(By.ID, 'userPw').send_keys('gksk1234!@#$')
driver.find_element(By.XPATH, '/html/body/div[2]/div[2]/form/div/div[2]/div/a').click()
driver.implicitly_wait(15)

# 예악페이지 들어가기 ( 1회차 예약 )
driver.get('https://gdsports.or.kr/info/soccer.do')
driver.find_element(By.XPATH, '/html/body/div[2]/div[2]/div[2]/div[1]/ul/li[4]/a').click()
driver.find_element(By.XPATH, '/html/body/div[2]/form/div[1]/div[2]/div[2]/div[3]/div/div[1]/div/div/a[2]').click()

# 예약 시작 
#driver.find_element(By.XPATH, '//*[@id="20220806"]').click() # id="날짜" 특정날짜를 알수있음
driver.find_element(By.XPATH, '//*[@id="20220820"]').click() # id="날짜" 특정날짜를 알수있음
driver.find_element(By.XPATH, '/html/body/div[2]/form/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[3]/table/tbody/tr[7]/td[1]/div/label/span').click()
driver.find_element(By.XPATH, '//*[@id="inputButton"]').click()
Alert(driver).accept()
time.sleep(2)
Alert(driver).accept()
driver.implicitly_wait(15)   

# 예약페이지

# /html/body/div[2]/form/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[3]/table/tbody/tr[5]/td[1]/div/label/span
# /html/body/div[2]/form/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[3]/table/tbody/tr[1]/td[1]/div/label/span (8시~10시)
# /html/body/div[2]/form/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[3]/table/tbody/tr[2]/td[1]
# /html/body/div[2]/form/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[3]/table/tbody/tr[7]/td[1]/div/label/span (20시~22시)