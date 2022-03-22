
# Markov-Chain 矩阵乘法的理解

Markov-Chain揭示了以概率为转移条件的状态转移规律

## 矩阵的理解角度
- 矩阵的项是一个2对1的映射

- 矩阵行/列是一组1对1的映射

- 矩阵行/列是一组n维向量

$$a_{ij} : f(i,j) \rightarrow a$$

## 矩阵乘法的理解角度
对n维向量的每个值做映射，再求和
$$
\begin{bmatrix}
    a_1  \cdots a_n
\end{bmatrix}
\begin{bmatrix}
    x_1 \\
    \vdots\\
    x_n
\end{bmatrix}
$$

## Markov-Chain 矩阵乘法理解角度

### 矩阵乘法理解的重点

- 对n维向量映射的理解,即$a_ix_i$
- 对每个维度向量映射后求和的理解,即 $\sum_{n=1}^N$
$$
\begin{bmatrix}
    a_{11}& \cdots & a_{1n}\\
    \vdots&  \ddots & \vdots\\
    a_{n1} & \cdots & a_{nn}
\end{bmatrix}
\begin{bmatrix}
    x_{11}& \cdots & x_{1n}\\
    \vdots&  \ddots & \vdots\\
    x_{n1} & \cdots & x_{nn}
\end{bmatrix}
=
\begin{bmatrix}
    \sum_{n=1}^N a_{1i}x_{i1}& \cdots & \sum_{n=1}^N a_{1i}x_{in}\\
    \vdots&  \ddots & \vdots\\
    \sum_{n=1}^N a_{ni}x_{in}& \cdots & \sum_{n=1}^N a_{ni}x_{in}\\
\end{bmatrix}

=
y
$$

$a_{ij}$ 和 $x_{ij}$ 的含义: 从状态i转移到状态j的概率

$a_{ij}x_i$ 的映射含义:  
