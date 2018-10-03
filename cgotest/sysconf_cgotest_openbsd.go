// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_cgotest

/*
#include <unistd.h>
*/
import "C"

import (
	"testing"

	"github.com/tklauser/go-sysconf"
)

func testSysconfCgoMatch(t *testing.T) {
	testCases := []testCase{
		{sysconf.SC_AIO_LISTIO_MAX, C._SC_AIO_LISTIO_MAX, "AIO_LISTIO_MAX"},
		{sysconf.SC_AIO_MAX, C._SC_AIO_MAX, "AIO_MAX"},
		{sysconf.SC_AIO_PRIO_DELTA_MAX, C._SC_AIO_PRIO_DELTA_MAX, "AIO_PRIO_DELTA_MAX"},
		{sysconf.SC_ARG_MAX, C._SC_ARG_MAX, "ARG_MAX"},
		{sysconf.SC_ATEXIT_MAX, C._SC_ATEXIT_MAX, "ATEXIT_MAX"},
		{sysconf.SC_BC_BASE_MAX, C._SC_BC_BASE_MAX, "BC_BASE_MAX"},
		{sysconf.SC_BC_DIM_MAX, C._SC_BC_DIM_MAX, "BC_DIM_MAX"},
		{sysconf.SC_BC_SCALE_MAX, C._SC_BC_SCALE_MAX, "BC_SCALE_MAX"},
		{sysconf.SC_BC_STRING_MAX, C._SC_BC_STRING_MAX, "BC_STRING_MAX"},
		{sysconf.SC_CHILD_MAX, C._SC_CHILD_MAX, "CHILD_MAX"},
		{sysconf.SC_CLK_TCK, C._SC_CLK_TCK, "CLK_TCK"},
		{sysconf.SC_COLL_WEIGHTS_MAX, C._SC_COLL_WEIGHTS_MAX, "COLL_WEIGHTS_MAX"},
		{sysconf.SC_DELAYTIMER_MAX, C._SC_DELAYTIMER_MAX, "DELAYTIMER_MAX"},
		{sysconf.SC_EXPR_NEST_MAX, C._SC_EXPR_NEST_MAX, "EXPR_NEST_MAX"},
		{sysconf.SC_GETGR_R_SIZE_MAX, C._SC_GETGR_R_SIZE_MAX, "GETGR_R_SIZE_MAX"},
		{sysconf.SC_GETPW_R_SIZE_MAX, C._SC_GETPW_R_SIZE_MAX, "GETPW_R_SIZE_MAX"},
		{sysconf.SC_HOST_NAME_MAX, C._SC_HOST_NAME_MAX, "HOST_NAME_MAX"},
		{sysconf.SC_IOV_MAX, C._SC_IOV_MAX, "IOV_MAX"},
		{sysconf.SC_LINE_MAX, C._SC_LINE_MAX, "LINE_MAX"},
		{sysconf.SC_LOGIN_NAME_MAX, C._SC_LOGIN_NAME_MAX, "LOGIN_NAME_MAX"},
		{sysconf.SC_NGROUPS_MAX, C._SC_NGROUPS_MAX, "NGROUPS_MAX"},
		{sysconf.SC_OPEN_MAX, C._SC_OPEN_MAX, "OPEN_MAX"},
		{sysconf.SC_PAGE_SIZE, C._SC_PAGE_SIZE, "PAGE_SIZE"},
		{sysconf.SC_PAGESIZE, C._SC_PAGESIZE, "PAGESIZE"},
		{sysconf.SC_THREAD_DESTRUCTOR_ITERATIONS, C._SC_THREAD_DESTRUCTOR_ITERATIONS, "PTHREAD_DESTRUCTOR_ITERATIONS"},
		{sysconf.SC_THREAD_KEYS_MAX, C._SC_THREAD_KEYS_MAX, "PTHREAD_KEYS_MAX"},
		{sysconf.SC_THREAD_STACK_MIN, C._SC_THREAD_STACK_MIN, "PTHREAD_STACK_MIN"},
		{sysconf.SC_THREAD_THREADS_MAX, C._SC_THREAD_THREADS_MAX, "PTHREAD_THREADS_MAX"},
		{sysconf.SC_RE_DUP_MAX, C._SC_RE_DUP_MAX, "RE_DUP_MAX"},
		{sysconf.SC_SEM_NSEMS_MAX, C._SC_SEM_NSEMS_MAX, "SEM_NSEMS_MAX"},
		{sysconf.SC_SEM_VALUE_MAX, C._SC_SEM_VALUE_MAX, "SEM_VALUE_MAX"},
		{sysconf.SC_SIGQUEUE_MAX, C._SC_SIGQUEUE_MAX, "SIGQUEUE_MAX"},
		{sysconf.SC_STREAM_MAX, C._SC_STREAM_MAX, "STREAM_MAX"},
		{sysconf.SC_SYMLOOP_MAX, C._SC_SYMLOOP_MAX, "SYMLOOP_MAX"},
		{sysconf.SC_TIMER_MAX, C._SC_TIMER_MAX, "TIMER_MAX"},
		{sysconf.SC_TTY_NAME_MAX, C._SC_TTY_NAME_MAX, "TTY_NAME_MAX"},
		{sysconf.SC_TZNAME_MAX, C._SC_TZNAME_MAX, "TZNAME_MAX"},

		{sysconf.SC_ADVISORY_INFO, C._SC_ADVISORY_INFO, "_POSIX_ADVISORY_INFO"},
		{sysconf.SC_ASYNCHRONOUS_IO, C._SC_ASYNCHRONOUS_IO, "_POSIX_ASYNCHRONOUS_IO"},
		{sysconf.SC_BARRIERS, C._SC_BARRIERS, "_POSIX_BARRIERS"},
		{sysconf.SC_CLOCK_SELECTION, C._SC_CLOCK_SELECTION, "_POSIX_CLOCK_SELECTION"},
		{sysconf.SC_CPUTIME, C._SC_CPUTIME, "_POSIX_CPUTIME"},
		{sysconf.SC_FSYNC, C._SC_FSYNC, "_POSIX_FSYNC"},
		{sysconf.SC_IPV6, C._SC_IPV6, "_POSIX_IPV6"},
		{sysconf.SC_JOB_CONTROL, C._SC_JOB_CONTROL, "_POSIX_JOB_CONTROL"},
		{sysconf.SC_MAPPED_FILES, C._SC_MAPPED_FILES, "_POSIX_MAPPED_FILES"},
		{sysconf.SC_MEMLOCK, C._SC_MEMLOCK, "_POSIX_MEMLOCK"},
		{sysconf.SC_MEMLOCK_RANGE, C._SC_MEMLOCK_RANGE, "_POSIX_MEMLOCK_RANGE"},
		{sysconf.SC_MEMORY_PROTECTION, C._SC_MEMORY_PROTECTION, "_POSIX_MEMORY_PROTECTION"},
		{sysconf.SC_MESSAGE_PASSING, C._SC_MESSAGE_PASSING, "_POSIX_MESSAGE_PASSING"},
		{sysconf.SC_MONOTONIC_CLOCK, C._SC_MONOTONIC_CLOCK, "_POSIX_MONOTONIC_CLOCK"},
		{sysconf.SC_PRIORITIZED_IO, C._SC_PRIORITIZED_IO, "_POSIX_PRIORITIZED_IO"},
		{sysconf.SC_PRIORITY_SCHEDULING, C._SC_PRIORITY_SCHEDULING, "_POSIX_PRIORITY_SCHEDULING"},
		{sysconf.SC_RAW_SOCKETS, C._SC_RAW_SOCKETS, "_POSIX_RAW_SOCKETS"},
		{sysconf.SC_READER_WRITER_LOCKS, C._SC_READER_WRITER_LOCKS, "_POSIX_READER_WRITER_LOCKS"},
		{sysconf.SC_REALTIME_SIGNALS, C._SC_REALTIME_SIGNALS, "_POSIX_REALTIME_SIGNALS"},
		{sysconf.SC_REGEXP, C._SC_REGEXP, "_POSIX_REGEXP"},
		{sysconf.SC_SAVED_IDS, C._SC_SAVED_IDS, "_POSIX_SAVED_IDS"},
		{sysconf.SC_SEMAPHORES, C._SC_SEMAPHORES, "_POSIX_SEMAPHORES"},
		{sysconf.SC_SHARED_MEMORY_OBJECTS, C._SC_SHARED_MEMORY_OBJECTS, "_POSIX_SHARED_MEMORY_OBJECTS"},
		{sysconf.SC_SHELL, C._SC_SHELL, "_POSIX_SHELL"},
		{sysconf.SC_SPAWN, C._SC_SPAWN, "_POSIX_SPAWN"},
		{sysconf.SC_SPIN_LOCKS, C._SC_SPIN_LOCKS, "_POSIX_SPIN_LOCKS"},
		{sysconf.SC_SPORADIC_SERVER, C._SC_SPORADIC_SERVER, "_POSIX_SPORADIC_SERVER"},
		{sysconf.SC_SYNCHRONIZED_IO, C._SC_SYNCHRONIZED_IO, "_POSIX_SYNCHRONIZED_IO"},
		{sysconf.SC_THREAD_ATTR_STACKADDR, C._SC_THREAD_ATTR_STACKADDR, "_POSIX_THREAD_ATTR_STACKADDR"},
		{sysconf.SC_THREAD_ATTR_STACKSIZE, C._SC_THREAD_ATTR_STACKSIZE, "_POSIX_THREAD_ATTR_STACKSIZE"},
		{sysconf.SC_THREAD_CPUTIME, C._SC_THREAD_CPUTIME, "_POSIX_THREAD_CPUTIME"},
		{sysconf.SC_THREAD_PRIO_INHERIT, C._SC_THREAD_PRIO_INHERIT, "_POSIX_THREAD_PRIO_INHERIT"},
		{sysconf.SC_THREAD_PRIO_PROTECT, C._SC_THREAD_PRIO_PROTECT, "_POSIX_THREAD_PRIO_PROTECT"},
		{sysconf.SC_THREAD_PRIORITY_SCHEDULING, C._SC_THREAD_PRIORITY_SCHEDULING, "_POSIX_THREAD_PRIORITY_SCHEDULING"},
		{sysconf.SC_THREAD_PROCESS_SHARED, C._SC_THREAD_PROCESS_SHARED, "_POSIX_THREAD_PROCESS_SHARED"},
		{sysconf.SC_THREAD_SAFE_FUNCTIONS, C._SC_THREAD_SAFE_FUNCTIONS, "_POSIX_THREAD_SAFE_FUNCTIONS"},
		{sysconf.SC_THREAD_SPORADIC_SERVER, C._SC_THREAD_SPORADIC_SERVER, "_POSIX_THREAD_SPORADIC_SERVER"},
		{sysconf.SC_THREADS, C._SC_THREADS, "_POSIX_THREADS"},
		{sysconf.SC_TIMEOUTS, C._SC_TIMEOUTS, "_POSIX_TIMEOUTS"},
		{sysconf.SC_TIMERS, C._SC_TIMERS, "_POSIX_TIMERS"},
		{sysconf.SC_TRACE, C._SC_TRACE, "_POSIX_TRACE"},
		{sysconf.SC_TRACE_EVENT_FILTER, C._SC_TRACE_EVENT_FILTER, "_POSIX_TRACE_EVENT_FILTER"},
		{sysconf.SC_TRACE_EVENT_NAME_MAX, C._SC_TRACE_EVENT_NAME_MAX, "_POSIX_TRACE_EVENT_NAME_MAX"},
		{sysconf.SC_TRACE_INHERIT, C._SC_TRACE_INHERIT, "_POSIX_TRACE_INHERIT"},
		{sysconf.SC_TRACE_LOG, C._SC_TRACE_LOG, "_POSIX_TRACE_LOG"},
		{sysconf.SC_TYPED_MEMORY_OBJECTS, C._SC_TYPED_MEMORY_OBJECTS, "_POSIX_TYPED_MEMORY_OBJECTS"},
		{sysconf.SC_VERSION, C._SC_VERSION, "_POSIX_VERSION"},

		{sysconf.SC_V7_ILP32_OFF32, C._SC_V7_ILP32_OFF32, "_POSIX_V7_ILP32_OFF32"},
		/*
			{sysconf.SC_V7_ILP32_OFFBIG, C._SC_V7_ILP32_OFFBIG, "_POSIX_V7_ILP32_OFFBIG"},
			{sysconf.SC_V7_LP64_OFF64, C._SC_V7_LP64_OFF64, "_POSIX_V7_LP64_OFF64"},
			{sysconf.SC_V7_LPBIG_OFFBIG, C._SC_V7_LPBIG_OFFBIG, "_POSIX_V7_LPBIG_OFFBIG"},
		*/

		{sysconf.SC_V6_ILP32_OFF32, C._SC_V6_ILP32_OFF32, "_POSIX_V6_ILP32_OFF32"},
		/*
			{sysconf.SC_V6_ILP32_OFFBIG, C._SC_V6_ILP32_OFFBIG, "_POSIX_V6_ILP32_OFFBIG"},
			{sysconf.SC_V6_LP64_OFF64, C._SC_V6_LP64_OFF64, "_POSIX_V6_LP64_OFF64"},
			{sysconf.SC_V6_LPBIG_OFFBIG, C._SC_V6_LPBIG_OFFBIG, "_POSIX_V6_LPBIG_OFFBIG"},
		*/

		{sysconf.SC_2_C_BIND, C._SC_2_C_BIND, "_POSIX2_C_BIND"},
		{sysconf.SC_2_C_DEV, C._SC_2_C_DEV, "_POSIX2_C_DEV"},
		{sysconf.SC_2_CHAR_TERM, C._SC_2_CHAR_TERM, "_POSIX2_CHAR_TERM"},
		{sysconf.SC_2_FORT_DEV, C._SC_2_FORT_DEV, "_POSIX2_FORT_DEV"},
		{sysconf.SC_2_FORT_RUN, C._SC_2_FORT_RUN, "_POSIX2_FORT_RUN"},
		{sysconf.SC_2_LOCALEDEF, C._SC_2_LOCALEDEF, "_POSIX2_LOCALEDEF"},
		{sysconf.SC_2_PBS, C._SC_2_PBS, "_POSIX2_PBS"},
		{sysconf.SC_2_PBS_ACCOUNTING, C._SC_2_PBS_ACCOUNTING, "_POSIX2_PBS_ACCOUNTING"},
		{sysconf.SC_2_PBS_CHECKPOINT, C._SC_2_PBS_CHECKPOINT, "_POSIX2_PBS_CHECKPOINT"},
		{sysconf.SC_2_PBS_LOCATE, C._SC_2_PBS_LOCATE, "_POSIX2_PBS_LOCATE"},
		{sysconf.SC_2_PBS_MESSAGE, C._SC_2_PBS_MESSAGE, "_POSIX2_PBS_MESSAGE"},
		{sysconf.SC_2_PBS_TRACK, C._SC_2_PBS_TRACK, "_POSIX2_PBS_TRACK"},
		{sysconf.SC_2_SW_DEV, C._SC_2_SW_DEV, "_POSIX2_SW_DEV"},
		{sysconf.SC_2_UPE, C._SC_2_UPE, "_POSIX2_UPE"},
		{sysconf.SC_2_VERSION, C._SC_2_VERSION, "_POSIX2_VERSION"},

		{sysconf.SC_XOPEN_CRYPT, C._SC_XOPEN_CRYPT, "_XOPEN_CRYPT"},
		{sysconf.SC_XOPEN_ENH_I18N, C._SC_XOPEN_ENH_I18N, "_XOPEN_ENH_I18N"},
		{sysconf.SC_XOPEN_REALTIME, C._SC_XOPEN_REALTIME, "_XOPEN_REALTIME"},
		{sysconf.SC_XOPEN_REALTIME_THREADS, C._SC_XOPEN_REALTIME_THREADS, "_XOPEN_REALTIME_THREADS"},
		{sysconf.SC_XOPEN_SHM, C._SC_XOPEN_SHM, "_XOPEN_SHM"},
		{sysconf.SC_XOPEN_STREAMS, C._SC_XOPEN_STREAMS, "_XOPEN_STREAMS"},
		{sysconf.SC_XOPEN_UNIX, C._SC_XOPEN_UNIX, "_XOPEN_UNIX"},
		{sysconf.SC_XOPEN_UUCP, C._SC_XOPEN_UUCP, "_XOPEN_UUCP"},

		// non-standard
		{sysconf.SC_PHYS_PAGES, C._SC_PHYS_PAGES, "_PHYS_PAGES"},
		// AV_PHYS_PAGES might change between calling Go and C version
		// of sysconf. Don't test it for now.
		{sysconf.SC_NPROCESSORS_CONF, C._SC_NPROCESSORS_CONF, "_NPROCESSORS_CONF"},
		{sysconf.SC_NPROCESSORS_ONLN, C._SC_NPROCESSORS_ONLN, "_NPROCESSORS_ONLN"},
	}
	for _, tc := range testCases {
		testSysconfGoCgo(t, tc)
	}

	testCasesInvalid := []testCase{
		{sysconf.SC_MQ_OPEN_MAX, C._SC_MQ_OPEN_MAX, "MQ_OPEN_MAX"},
		{sysconf.SC_MQ_PRIO_MAX, C._SC_MQ_PRIO_MAX, "MQ_PRIO_MAX"},
		{sysconf.SC_SS_REPL_MAX, C._SC_SS_REPL_MAX, "_POSIX_SS_REPL_MAX"},
		{sysconf.SC_TRACE_NAME_MAX, C._SC_TRACE_NAME_MAX, "_POSIX_TRACE_NAME_MAX"},
		{sysconf.SC_TRACE_SYS_MAX, C._SC_TRACE_SYS_MAX, "_POSIX_TRACE_SYS_MAX"},
		{sysconf.SC_TRACE_USER_EVENT_MAX, C._SC_TRACE_USER_EVENT_MAX, "_POSIX_TRACE_USER_EVENT_MAX"},
		{sysconf.SC_XOPEN_VERSION, C._SC_XOPEN_VERSION, "_XOPEN_VERSION"},
	}
	for _, tc := range testCasesInvalid {
		testSysconfGoCgoInvalid(t, tc)
	}
}
