class Solution(object):
    def gameOfLife(self, board):
        M, N = len(board[0]), len(board)
        cells_ind = [(x, y) for x in range(N) for y in range(M)]
        next_board = dict(zip(cells_ind, [0] * M * N))
    
        x_mods = [1,  1,  1, -1, -1, -1,  0,  0]
        y_mods = [1, -1,  0,  1, -1,  0,  1, -1] 
        coor_mods = []
            
        for i in range(8):
            coor_mods.append((x_mods[i], y_mods[i]))
        
        for x, y in cells_ind:
            adj_living_cells = 0    
            adj = 0
            
            for x_mod, y_mod in coor_mods:
                cur_x, cur_y = x + x_mod, y + y_mod
                if (cur_x, cur_y) in cells_ind:
                    adj += 1
                    if board[cur_x][cur_y] == 1:
                        adj_living_cells += 1
                        
            if board[x][y] == 1:
                if adj_living_cells < 2:
                    next_board[(x, y)] = 0
                elif adj_living_cells > 3:
                    next_board[(x, y)] = 0
                else:
                    next_board[(x, y)] = 1
            elif adj_living_cells == 3:
                next_board[(x, y)] = 1
                
        for (x, y), status in next_board.items():
            board[x][y] = status
