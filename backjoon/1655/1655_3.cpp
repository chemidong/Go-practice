//Problem 1655 in C++

#include <iostream>
#include <tuple>

typedef std::tuple<int, int> Childs;
const bool MAX = true;
const bool MIN = false;
int max[50001];
int n_max = -1;
int min[50001];
int n_min = -1;


int find(bool mode) {
    if(mode == MAX) return max[0];
    else return min[0];
}

int median() {
    if(n_max < n_min) return find(MIN);
    else return find(MAX);
}

int parent(int i) {
    return (i - 1) / 2;
}

Childs child(int i) {
    return std::make_tuple(i*2 + 1, i*2 + 2);
}

void change(int *a, int *b) {
    int t = *a;
    *a = *b;
    *b = t;
}

void insert(int v, bool mode) {
    int i, p;
    if(mode == MAX) {
        n_max++;
        if(n_max == 0) max[0] = v;
        else {
            i = n_max;
            max[i] = v;
            while(i > 0) {
                p = parent(i);
                if(max[p] < max[i]) {
                    change(&max[p], &max[i]);
                    i = p;
                } 
                else break;
            }
        }
    }else {
        n_min++;
        if(n_min == 0) min[0] = v;
        else {
            i = n_min;
            min[i] = v;
            while(i > 0) {
                p = parent(i);
                if(min[p] > min[i]) {
                    change(&min[p], &min[i]);
                    i = p;
                }
                else break;
            }
        }
    }
}
int remove(bool mode) {
    int ret;
    if (mode == MAX){
        ret = max[0];
        max[0] = max[n_max];
        n_max--;
        int i = 0;
        while (true) {
            Childs c = child(i);
            int left = std::get<0>(c);
            int right = std::get<1>(c);

            if (right <= n_max) {
                if (max[left] > max[right]) {
                    if (max[left] > max[i]) {
                        change(&max[i], &max[left]);
                        i = left;
                    } 
                    else break;
                } else {
                    if (max[right] > max[i]) {
                        change(&max[i], &max[right]);
                        i = right;
                    } 
                    else break;
                }
            } else if (left == n_max) {
                if (max[left] > max[i]) change(&max[i], &max[left]);
                break;
            } 
            else break;
        }
    } else {
        ret = min[0];
        min[0] = min[n_min];
        n_min--;
        int i = 0;
        while (true) {
            Childs c = child(i);
            int left = std::get<0>(c);
            int right = std::get<1>(c);

            if (right <= n_min) {
                if (min[left] < min[right]) {
                    if (min[left] < min[i]) {
                        change(&min[i], &min[left]);
                        i = left;
                    } 
                    else break;
                } else {
                    if (min[right] < min[i]) {
                        change(&min[i], &min[right]);
                        i = right;
                    } 
                    else break;
                }
            } else if (left == n_min) {
                if (min[left] < min[i]) change(&min[i], &min[left]);
                break;
            } 
            else break;
        }
    }
    return ret;
}

int main() {
    std::ios::sync_with_stdio(false);
    std::cin.tie(0);
    int N, v;
    int med = -10001;
    std::cin >> N;
    for (int i = 0; i < N; i++){
        std::cin >> v;
        if(med < v) insert(v, MIN);
        else insert(v, MAX);
        if(n_max > n_min && n_max - n_min == 2){
            v = remove(MAX);
            insert(v, MIN);
        } else if(n_max < n_min && n_min - n_max == 2){
            v = remove(MIN);
            insert(v, MAX);
        }
        med = median();
        std::cout << med << '\n';
    }
    return 0;
}