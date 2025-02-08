from helpy import helpy
def app():
    print("------- Hello welcome to Console Unet ------")
    print("For help enter help")
    run = True
    while run:
        command = input("Enter command")
        if command == "help":
            helpy()
