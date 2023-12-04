#include <iostream>
#include <stdio.h>
#include <iomanip>
#include <string>
#include <cstdlib>
using namespace std;
// 预先开辟空间，实际数据从data_in.txt读取
#define MAXVEX 100  // 顶点数
#define MAXEDGE 500 // 边数

#define MAX 65535 // 无边时的权值

typedef struct
{
    string vexs[MAXVEX];
    int arc[MAXVEX][MAXVEX];
    int numVertexes, numEdges;
} MGraph;

typedef int PathArc[MAXVEX];
typedef int ShortPathTable[MAXVEX];

/**********************函数声明  Start***********************************/
// 读取data_in.txt，建立邻接矩阵
void CreateMGraph(MGraph *G);
// Dijkstra算法，求有向网G的pos顶点到其余顶点v的最短路径P[v]及带权长度D[v]
void ShortestPath_Dijkstra(MGraph MG, int pos, PathArc P, ShortPathTable D);
// 将倒序查找的路径节点，顺序输出
void searchPath(MGraph MG, int pos, PathArc P, ShortPathTable D);
// 查找任意两点之间最短路径及权值
void twoPointPath_Dijkstra(int posStart, int posEnd, MGraph MG, PathArc P, ShortPathTable D);
void printGraph(MGraph G);         // 打印邻接矩阵
int LocateVex(MGraph G, string u); // 返回顶点u在图中的位置
void redEdgesMod(MGraph &G);       // 修改红色边权值为65535
void menu();                       // 主菜单
/**********************函数声明  End***********************************/

int main(void)
{
    PathArc P;
    ShortPathTable D;

    MGraph MG = {
        {"S", "N1", "N2", "N3", "N4",
         "N5", "N6", "N7", "N8", "N9",
         "N10", "N11", "N12", "N13", "N14",
         "N15", "N16", "E"},
        // 权值无限大,代表两边不相连，
        // 邻接矩阵 18*18
        {
            0},
        MAXVEX, // 顶点数
        MAXEDGE // 边数
    };
    int f;
    while (true)
    {
        cout << endl;
        menu();
        cin >> f;
        if (f == 1)
        {
            CreateMGraph(&MG); // 读取数据文件，顶点数+边数+邻接矩阵
            printGraph(MG);
        }
        else if (f == 2)
        {
            redEdgesMod(MG);
            // printGraph(MG);
            ShortestPath_Dijkstra(MG, 0, P, D);
            searchPath(MG, 0, P, D);
        }
        else if (f == 3)
        {
            cout << "经过绿点且不经过红边" << endl;
            twoPointPath_Dijkstra(0, 7, MG, P, D);
            twoPointPath_Dijkstra(7, 12, MG, P, D);
            twoPointPath_Dijkstra(12, 17, MG, P, D);
            cout << "S->N3->N7->N6->N12->N16->E.   "
                 << "权值为7,共计7个节点" << endl;
            //	twoPointPath_Dijkstra(0,12,MG, P, D);
            //	twoPointPath_Dijkstra(12,7,MG, P, D);
            //	twoPointPath_Dijkstra(7,17,MG, P, D);
        }
        else if (f == 4)
        {
            cout << "经过绿边，尽可能多的经过绿点，但不经过红边" << endl;
            twoPointPath_Dijkstra(0, 2, MG, P, D);
            twoPointPath_Dijkstra(4, 14, MG, P, D);
            twoPointPath_Dijkstra(13, 17, MG, P, D);
            cout << "S->N2->N4->N5->N6->N14-N13->E. "
                 << "权值为10，共计8个节点" << endl;
        }
        else if (f == 5)
        {
            cout << ">>>退出成功！" << endl;
            exit(0);
        }
    }
}

/* 构建图 */
void CreateMGraph(MGraph *G)
{
    FILE *pfile = NULL;
    int i = 0;
    int j = 0;
    int weight = 0;
    int k = 0;

    // 从文件中读入数据
    pfile = fopen("data_in.txt", "r");
    if (pfile == NULL)
    {
        printf("ERROR:The file 'data_in.txt' was not opened\n");
        exit(0);
    }
    // 读取顶点数和边数
    fscanf(pfile, "%d  %d", &(G->numVertexes), &(G->numEdges));
    // printf("%d  %d\n",G->numVertexes,G->numEdges);

    // 初始化邻接矩阵
    for (i = 0; i < G->numVertexes; i++)
    {
        for (j = 0; j < G->numVertexes; j++)
        {
            G->arc[i][j] = MAX; //
        }
    }
    // 读取边和权值信息
    for (k = 0; k < G->numEdges; k++) // 从文件读入各城市起点，终点和权重
    {
        fscanf(pfile, "%d %d %d", &i, &j, &weight);
        G->arc[i][j] = weight;
        // G.edges[][city2]=weight;
    }
    // 对称阵
    for (i = 0; i < G->numVertexes; i++) // G->numVertexes
    {
        for (j = i; j < G->numVertexes; j++) // G->numVertexes
        {
            G->arc[j][i] = G->arc[i][j];
        }
    }
}

/*  Dijkstra算法，求有向网G的pos顶点到其余顶点v的最短路径P[v]及带权长度D[v] */
/*  P[v]的值为前驱顶点下标,D[v]表示pos到v的最短路径长度和 */
/*  pos 取值 0～MG.numVertexs-1 */
void ShortestPath_Dijkstra(MGraph MG, int pos, PathArc P, ShortPathTable D)
{
    int v, w, k, min;
    int final[MAXVEX]; /* final[w]=1表示求得顶点pos至w的最短路径 */
    for (v = 0; v < MG.numVertexes; v++)
    {
        final[v] = 0;          /* 全部顶点初始化为未知最短路径状态 */
        D[v] = MG.arc[pos][v]; /* 将与pos点有连线的顶点加上权值,可达为实际权值，不可达为MAX */
        P[v] = 0;              /* 初始化路径数组P为0  */
    }

    D[pos] = 0;     /*说明源点pos没有到自身的路径 */
    P[pos] = -1;    /* -1表示自身无前驱顶点*/
    final[pos] = 1; /* pos至pos不需要求路径 */
    /* 开始主循环，每次求得pos到某个v顶点的最短路径 */
    for (v = 1; v < MG.numVertexes; v++)
    {
        min = MAX;                           /* 当前所知离pos顶点的最近距离 */
        for (w = 0; w < MG.numVertexes; w++) /* 寻找离pos最近的顶点 */
        {
            if (!final[w] && D[w] < min)
            {
                k = w;
                min = D[w]; /* w顶点离pos顶点更近 */
            }
        }
        final[k] = 1;                        /* 将目前找到的最近的顶点置为1 */
        for (w = 0; w < MG.numVertexes; w++) /* 修正当前最短路径及距离 */
        {
            if (!final[w] && (min + MG.arc[k][w] < D[w]))
            {
                /*  说明找到了更短的路径，修改D[w]和P[w] */
                D[w] = min + MG.arc[k][w]; /* 修改当前路径长度 */
                P[w] = k;
            }
        }
    }
    /* 结束循环，若P[w] = 0;说明顶点w的前驱为pos */
}
// 利用D计算源节点到每个节点最小路径
// 利用P计算源节点到每个节点路径节点情况
// 利用公共全局变量： MAXVEX， 形参：MG P,D
// 将倒序查找的路径节点，顺序输出
void searchPath(MGraph MG, int pos, PathArc P, ShortPathTable D)
{
    int i, j;
    int n = 0;                       // pathTemp的计数器
    int realMAXVEX = MG.numVertexes; // 实际顶点数
    int pathTemp[MAXVEX] = {0};      // 依次记住前驱点，之后倒序输出，最后一个节点为源节点

    cout << MG.vexs[pos] << "到各顶点的最短路径及长度为：" << endl; // 修改为1，即标号2
    for (i = 0; i < realMAXVEX; i++)
    {
        j = i; // 不能删除，因为j在下面循环中变化

        n = 0; // 为下一次做准备
        // 清空路径节点记录，为下一次做准备
        for (int m = 0; m < realMAXVEX; m++)
        {
            pathTemp[m] = 0;
        }

        // 路径是否可达（若可达则倒序查找节点)
        if (D[j] != MAX)
        {
            cout << MG.vexs[pos] << "-" << MG.vexs[j] << "的最短路径的权值为：" << D[j];
            while (P[j] != -1 && P[j] != 0)
            {

                pathTemp[n] = j; // 记住j开始的每一个前驱点(除去前驱点为源节点)
                n++;
                // 倒序查找j->P[j]->j(坐标值为新的下标)
                j = P[j];
            }

            pathTemp[n] = j; // 找到前驱点的值为0，表示此时前驱点为源点
            n++;
            pathTemp[n] = pos; // 数组最后一个节点为源节点
        }
        else if (D[j] == MAX)
        {
            cout << MG.vexs[pos] << "-" << MG.vexs[j] << ":"
                 << "不可达" << endl;
        }
        // 打印最短路径
        cout << "   最短路径为：";
        for (; n >= 0; n--)
        {
            // 不考虑最短路径为0（源点到源点 )
            if (P[j] > -1)
            {
                cout << MG.vexs[pathTemp[n]] << " ";
            }
        }
        cout << endl;
    }
}
void twoPointPath_Dijkstra(int posStart, int posEnd, MGraph MG, PathArc P, ShortPathTable D)
{
    ShortestPath_Dijkstra(MG, posStart, P, D);
    cout << MG.vexs[posStart] << "-" << MG.vexs[posEnd] << "段之间，最短路径的权值为：" << D[posEnd] << endl;
    int j = posEnd;
    while (P[j] != -1 && P[j] != 0)
    {
        cout << MG.vexs[j] << "<-" << MG.vexs[P[j]] << "  ";
        j = P[j];
    }
    cout << MG.vexs[j] << "<-" << MG.vexs[posStart] << "  ";
    cout << endl;
}
void printGraph(MGraph G) // 打印邻接矩阵
{

    // 输出邻接矩阵信息
    cout << "<---无向网信息---->" << endl;
    cout << "顶点数和边数：" << G.numVertexes << "  " << G.numEdges << endl;
    cout << "邻接矩阵：" << endl;
    for (int m = 0; m < G.numVertexes; m++)
    {
        for (int n = 0; n < G.numVertexes; n++)
        {
            // cout <<setw(2)<<G.arc[m][n] << "  " ;
            printf("%5d ", G.arc[m][n]);
        }
        cout << endl; // 每输出一行后进行换行
    }
}
int LocateVex(MGraph G, string u) // 返回顶点u在图中的位置
{
    for (int i = 0; i < G.numVertexes; i++)
        if (G.vexs[i] == u)
            return i;
    return -1;
}
// 修改红色边权值为65535
void redEdgesMod(MGraph &G) // 形参传址
{
    /*    string v1, v2;
        int w;
        int radEdges;

        int i, j, k;
        cout<<"请输入修改红色边的条数 "<<endl;
        cin>>radEdges;
        cout<<"请输入红色边对应的顶点和修改权值："<<endl;
        for(k=0; k<radEdges; k++)
        {
            cin>>v1>>v2>>w;
            i=LocateVex(G, v1);
            j=LocateVex(G, v2);
            cout<<i<<j<<endl;
            G.arc[i][j]=w;
            G.arc[j][i]=w;  //矩阵为对称矩阵
        }
    */
    // 修改N11-N12的权值为MAX
    string v1, v2;
    int i, j;
    v1 = "N11";
    v2 = "N12";
    i = LocateVex(G, v1);
    j = LocateVex(G, v2);
    G.arc[i][j] = MAX;
    G.arc[j][i] = MAX; // 矩阵为对称矩阵
    cout << "红色边(N11-N12),修改权值为MAX=65535" << endl;
}
// 主菜单
void menu()
{
    cout << "|*********************************************************************|" << endl;
    cout << "|-----------------------------小蚂蚁找食物程序------------------------|" << endl;
    cout << "|*********************************************************************|" << endl;
    cout << "|-----------------------1-读取数据文并显示邻接矩阵--------------------|" << endl;
    cout << "|-----------------------2-不走红边并求最短路径查询--------------------|" << endl;
    cout << "|-----------------------3-经过绿点，且不走红边最短路径----------------|" << endl;
    cout << "|-----------------------4-经过绿边，且不走红边最短路径----------------|" << endl;
    cout << "|-----------------------5-退出小蚂蚁找食物程序------------------------|" << endl;
    cout << "|*********************************************************************|" << endl;
    cout << ">>>请选择：";
}
