#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#define N 1000
/*定义复数类型*/
typedef struct
{
	double real;
	double img;
} complex;
complex x[N], *W;
/*输入序列,变换核*/
int Nx = 32;
;
/* 输入源数据个数*/
int NFFT = (int)pow(2.0, ceil(log((double)Nx) / log(2.0)));
/*输入序列的大小，在本程序中仅限2的次幂*/
double table[32] = {
	0};
double PI;
/*圆周率*/
void fft();
/*快速傅里叶变换*/
void initW();
/*初始化变换核*/
void change();
/*变址*/
void add(complex, complex, complex *);
/*复数加法*/
void mul(complex, complex, complex *);
/*复数乘法*/
void sub(complex, complex, complex *);
/*复数减法*/
void output();
/*输出快速傅里叶变换的结果*/
/************根据源数据个数计算FFT点数********************************/
/*输入源数据个数 */
/* 计算FFT的点数：as the next higer power of 2 >=Nx*/
/* 使用额外空间进行数组旋转 */
void RotateArrayK_Ext(double nArray[], int nCount, int k)
{
	if (nArray == NULL || nCount <= 0 || k < 0 || k > nCount)
		return;
	double *pArrayTemp = new double[nCount];
	if (NULL == pArrayTemp)
		return;
	/* 备份数据 */
	for (int i = 0; i < nCount; ++i)
	{
		pArrayTemp[i] = nArray[i];
	}
	/* 进行错位复制 */
	int i = 0;
	for (i = 0; i < k; ++i)
	{
		nArray[i] = pArrayTemp[nCount - k + i];
	}
	for (int j = 0; i < nCount; ++i)
	{
		nArray[i] = pArrayTemp[j++];
	}
}
void jisuan()
{
	printf("The result are as follows：\n");
	/* 求复数模值 */
	for (int i = 0; i < NFFT; i++)
	{
		/* table[i]=(sqrt(x[i].real*x[i].real+x[i].img*x[i].img)); ///(i==0?NFFT:NFFT/2)); */
		table[i] = (sqrt(x[i].real * x[i].real + x[i].img * x[i].img) / (i == 0 ? NFFT : NFFT / 2));
		/*	printf("%.4f \n",table[i]); */
	}
	/* fftshift */
	RotateArrayK_Ext(table, NFFT, NFFT / 2);
	for (int i = 0; i < NFFT; i++)
	{
		printf("%d   %.4f \n", i, table[i]);
	}
}
int main()
{
	int i;
	/*输出结果*/
	double sumCos, averageCos;
	/* 去除cos直流分量 */
	double sumSin, averageSin;
	/* 去除Sin直流分量 */
	system("cls");
	PI = atan(1) * 4;
	printf("  数据点数&&FFT点数     \n");
	printf("Nx = %d\n", Nx);
	printf("NFFT = %d \n", NFFT);
	/************源数据初始化********************************/
	/* input test data
	 *
	 * for(i=0;i<Nx;i++)
	 *
	 * {
	 *
	 * x[i].real=i;
	 *
	 * x[i].img=0.0;
	 *
	 * }
	 *
	 * for(i=Nx;i<NFFT;i++)
	 *
	 * {
	 *
	 * x[i].real=0.0;
	 *
	 * x[i].img=0.0;
	 *
	 * }*/
	x[0].real = 255;
	x[1].real = 232;
	x[2].real = 168;
	x[3].real = 88;
	x[4].real = 24;
	x[5].real = 0;
	x[6].real = 24;
	x[7].real = 88;
	x[8].real = 168;
	x[9].real = 232;
	x[10].real = 255;
	x[11].real = 232;
	x[12].real = 168;
	x[13].real = 88;
	x[14].real = 24;
	x[15].real = 0;
	x[16].real = 24;
	x[17].real = 88;
	x[18].real = 168;
	x[19].real = 232;
	x[20].real = 255;
	x[21].real = 232;
	x[22].real = 168;
	x[23].real = 88;
	x[24].real = 24;
	x[25].real = 0;
	x[26].real = 24;
	x[27].real = 88;
	x[28].real = 168;
	x[29].real = 232;
	x[30].real = 255;
	x[31].real = 232;
	x[0].img = 128;
	x[1].img = 203;
	x[2].img = 250;
	x[3].img = 250;
	x[4].img = 203;
	x[5].img = 128;
	x[6].img = 53;
	x[7].img = 6;
	x[8].img = 6;
	x[9].img = 53;
	x[10].img = 128;
	x[11].img = 203;
	x[12].img = 250;
	x[13].img = 250;
	x[14].img = 203;
	x[15].img = 128;
	x[16].img = 53;
	x[17].img = 6;
	x[18].img = 6;
	x[19].img = 53;
	x[20].img = 128;
	x[21].img = 203;
	x[22].img = 250;
	x[23].img = 250;
	x[24].img = 203;
	x[25].img = 128;
	x[26].img = 53;
	x[27].img = 6;
	x[28].img = 6;
	x[29].img = 53;
	x[30].img = 128;
	x[31].img = 203;
	/*分别去除实部cos、虚部sin直流分量*/
	for (i = 0; i < NFFT; i++)
	{
		sumCos += x[i].real;
		sumSin += x[i].img;
	}
	averageCos = sumCos / NFFT;
	averageSin = sumSin / NFFT;
	printf("sumCos =%f   averageCos = %f\n", sumCos, averageCos);
	printf("sumSin =%f   averageSin = %f\n", sumSin, averageSin);
	for (i = 0; i < NFFT; i++)
	{
		x[i].real -= averageCos;
		x[i].img -= averageSin;
	}
	/************源数据初始化(未使用，先用全局变量做)********************************/
	/* NFFT运算分配动态内存*/
	/* x=(complex*)malloc(NFFT*sizeof(complex)); */
	/************先倒序，后FFT计算********************************/
	printf("输出倒序后的序列\n");
	initW();
	/* 调用变换核 */
	fft();
	/* 调用快速傅里叶变换 */
	printf("输出FFT后的结果\n");
	output();
	/* 调用输出傅里叶变换结果函数 */
	/* 先求模值，再fftshift */
	jisuan();
	return (0);
}
/*快速傅里叶变换*/
void fft()
{
	int i = 0, j = 0, k = 0, l = 0;
	complex up, down, product;
	change();
	/* 调用变址函数 */
	for (i = 0; i < log(NFFT) / log(2); i++)
	/*一级蝶形运算 stage */ {
		l = 1 << i;
		for (j = 0; j < NFFT; j += 2 * l)
		/*一组蝶形运算 group,每个group的蝶形因子乘数不同*/ {
			for (k = 0; k < l; k++)
			/*一个蝶形运算 每个group内的蝶形运算*/ {
				mul(x[j + k + l], W[NFFT * k / 2 / l], &product);
				add(x[j + k], product, &up);
				sub(x[j + k], product, &down);
				x[j + k] = up;
				x[j + k + l] = down;
			}
		}
	}
}
/*初始化变换核，定义一个变换核，相当于旋转因子WAP*/
void initW()
{
	int i;
	W = (complex *)malloc(sizeof(complex) * NFFT);
	/* 生成变换核 */
	for (i = 0; i < NFFT; i++)
	{
		W[i].real = cos(2 * PI / NFFT * i);
		/* 用欧拉公式计算旋转因子 */
		W[i].img = -1 * sin(2 * PI / NFFT * i);
	}
}
/*变址计算，将x(n)码位倒置*/
void change()
{
	complex temp;
	unsigned short i = 0, j = 0, k = 0;
	double t;
	for (i = 0; i < NFFT; i++)
	{
		k = i;
		j = 0;
		t = (log(NFFT) / log(2));
		while ((t--) > 0)
		/* 利用按位与以及循环实现码位颠倒 */ {
			j = j << 1;
			j |= (k & 1);
			k = k >> 1;
		}
		if (j > i)
		/* 将x(n)的码位互换 */ {
			temp = x[i];
			x[i] = x[j];
			x[j] = temp;
		}
	}
	output();
}
/*输出傅里叶变换的结果*/
void output()
{
	int i;
	printf("The result are as follows：\n");
	for (i = 0; i < NFFT; i++)
	{
		printf("%.4f", x[i].real);
		printf("+ [%.4f] j\n", x[i].img);
		/*	if(x[i].img>=0.0001)printf("+%.4fj\n",x[i].img);
		 *
		 * else if(fabs(x[i].img)<0.0001)printf("\n");
		 *
		 * else printf("%.4fj\n",x[i].img);
		 *
		 */
	}
}
void add(complex a, complex b, complex *c)
/* 复数加法的定义 */ {
	c->real = a.real + b.real;
	c->img = a.img + b.img;
}
void mul(complex a, complex b, complex *c)
/* 复数乘法的定义 */ {
	c->real = a.real * b.real - a.img * b.img;
	c->img = a.real * b.img + a.img * b.real;
}
void sub(complex a, complex b, complex *c)
/* 复数减法的定义 */ {
	c->real = a.real - b.real;
	c->img = a.img - b.img;
}