import imageio
from PIL import ImageFont, Image, ImageDraw


# 制作60秒倒计时的背景图片
def make_picture(path):
    # 设置字体及字号
    font = ImageFont.truetype("simhei.ttf", 300)

    # 制作0——9的数字图片
    for idx in list([i for i in range(0, 10)]):
        # 生成一张白色背景图
        img = Image.new("RGB", (300, 300), (255, 255, 255))
        # 在背景图片上添加文字
        draw = ImageDraw.Draw(img)
        # 第一个为文本位置，第二个为文本内容，第三个为文本颜色，第四个为文本字体
        draw.text((80, 5), str(idx), (255, 0, 0), font)
        # 保存图片
        img.save(path + str(idx) + '.png')

    # 制作10-60的数字图片
    for idx in list([i for i in range(10, 61)]):
        # 生成一张白色背景图
        img = Image.new("RGB", (300, 300), (255, 255, 255))
        # 在背景图片上添加文字
        draw = ImageDraw.Draw(img)
        # 第一个为文本位置，第二个为文本内容，第三个为文本颜色，第四个为文本字体
        draw.text((0, 5), str(idx), (255, 0, 0), font)
        # 保存图片
        img.save(path + str(idx) + '.png')


# 制作gif图
def make_gif_imageio(path):
    list = [path + str(i) + '.png' for i in range(1, 61)]
    img_list = []
    for img_name in list:
        img_list.append(imageio.imread(img_name))
    img_list.reverse()
    # duration 切换秒数
    imageio.mimsave('60s.gif', img_list, 'GIF', duration=1)



path = "images/"

if __name__ == "__main__":
    make_picture(path)  # 制作图片
    make_gif_imageio(path)  # 制作gif图
