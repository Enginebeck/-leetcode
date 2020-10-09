package simple

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left := 1
	right := n
	for {
		rightIsBadVersion := isBadVersion(right)
		leftIsBadVersion := isBadVersion(left)
		middle := (right + left) / 2
		middleIsBadVersion := isBadVersion(middle)
		if leftIsBadVersion {
			return left
		} else if left == n {
			return 0
		} else if middleIsBadVersion {
			right = middle
		} else if rightIsBadVersion {
			if right == left+1 {
				return right
			} else {
				left = middle
			}
		}
	}
}

func isBadVersion(version int) bool {
	return version >= 4
}
