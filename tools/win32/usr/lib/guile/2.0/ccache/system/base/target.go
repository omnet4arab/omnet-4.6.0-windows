GOOF----LE-4-2.0�      ] ] 4 h      ] g  guile�	 �	g  define-module*�	 �	 �	g  system�	g  base�	g  target�		 �	
g  filenameS�	f  system/base/target.scm�	g  importsS�	g  rnrs�	g  bytevectors�	 �	 �	g  ice-9�	g  regex�	 �	 �	 �	g  exportsS�	g  target-type�	g  with-target�	g  
target-cpu�	g  target-vendor�	g  	target-os�	g  target-endianness�	g  target-word-size�	 �	g  set-current-module�	  �	! �	"g  foreign�	#" �	$g  sizeof�	%#$ �	&#$ �	'g  *�	(g  %native-word-size�	)g  
make-fluid�	*g  
%host-type�	+g  %target-type�	,g  native-endianness�	-g  %target-endianness�	.g  %target-word-size�	/g  string?�	0g  string-split�	1g  length�	2g  or-map�	3g  string-null?�	4g  error�	5f  invalid target�	6g  validate-target�	7g  triplet-cpu�	8g  cpu-endianness�	9g  triplet-pointer-size�	:g  string=?�	;g  string-match�	<f  
^i[0-9]86$�	=g  little�	>g  member�	?f  x86_64�	@f  ia64�	Af  	powerpcle�	Bf  powerpc64le�	Cf  mipsel�	Df  mips64el�	E?@ABCD �	Ff  sparc�	Gf  sparc64�	Hf  powerpc�	If  	powerpc64�	Jf  spu�	Kf  mips�	Lf  mips64�	MFGHIJKL �	Ng  big�	Of  ^arm.*el�	Pf  unknown CPU endianness�	Qf  ^mips64.*-gnuabi64�	Rf  ^mips64�	Sf  ^x86_64-.*-gnux32�	Tf  64$�	Uf  64[lbe][lbe]$�	VFHKC �	Wf  ^arm.*�	Xf  unknown CPU word size�	Yg  
triplet-os�	Zg  	substring�	[g  string-index�	\g  triplet-vendor�C 5    h  (  ]4	
5 4! >  "  G   4&'5(R4)i*i5+R4)i4,i5 5-R4)i(i5.R/012345   hX   �   ]4 5$  54 -545	�$  "  	45"  $   6C       �       g  target
		Q g  parts		> g  t		"	;  g  filenamef  system/base/target.scm�
	-
��		.	��		.	��		/	��		/	
��		0	��	"	0	��	"	0	��	0	1	��	C	.	
��	G	.	��	K	2	��	O	2	�� 		Q  g  nameg  validate-target� C6R67+-.89 hH   �   ]
4 >  "  G  4 5 454 5Y4>   ZCZF       �       g  target
		A g  thunk		A g  cpu			A  g  filenamef  system/base/target.scm�
	4
��		5	��		6	��		6	��	&	8	&��	-	9	%��	6	:	�� 		A	  g  nameg  with-target� CR:7*,;<=>EMNO4P    h`   �  ]4 455$  6 4 5$  C4 	5$  C4 
5$  C4 5$  C 6y      g  cpu
		`  g  filenamef  system/base/target.scm�
	<
��		>	��		>	��		>	��		>	��		?	��		@	��		@	��	!	@	��	%	@	��	'	A	��	)	B	��	/	B	��	1	B	��	5	@	��	7	D	��	9	E	��	?	E	��	A	E	��	E	@	��	G	G	��	I	H	��	M	H	��	Q	H	��	U	@	��	W	I	��	\	K	��	`	K	�� 		`  g  nameg  cpu-endianness�g  documentationf  Return the endianness for CPU.� C8R7;<QRSTU>VW4X:*Y(      h�   Y  ]	4 5"  �45$  	C4 5$  	C45$  	C4 5$  	C45$  	C45$  	C4	
5$  	C45$  	C64455$  44 5455$  C"��H"��D     Q      g  triplet
	 � g  cpu		 �  g  filenamef  system/base/target.scm�
	M
��		O	��			O	��		T	��		T	��		T	��		P	��	 	[	��	$	[	��	(	[	��	,	P	��	0	\	��	4	\	��	8	\	��	<	P	��	@	^	��	D	^	��	H	^	��	L	P	��	P	`	��	T	`	��	X	`	��	\	P	��	`	a	��	d	a	��	h	a	��	l	P	��	p	b	��	v	b	��	x	b	��	|	P	�� �	c	�� �	c	�� �	c	�� �	P	�� �	d	�� �	d	�� �	P	�� �	P	�� �	P	�� �	P	�� �	P	�� �	Q	�� �	Q	�� �	Q	/�� �	Q	�� �	P	�� 0	 �  g  nameg  triplet-pointer-size�g  documentationf  1Return the size of pointers in bytes for TRIPLET.� C9RZ[    h   �   ] 
4 -56      x       g  t
		  g  filenamef  system/base/target.scm�
	f
��		g	��		g	�� 		  g  nameg  triplet-cpu� C7R[Z     h(   �   ]	4 -5� 4 -56       �       g  t
		! g  start		!  g  filenamef  system/base/target.scm�
	i
��		j	��		j	��		j	��		k	��	!	k	�� 		!  g  nameg  triplet-vendor� C\R[Z        h    �   ]	4 -4 -5�5� 6�       g  t
		  g  start		   g  filenamef  system/base/target.scm�
	m
��		n	��	
	n	+��		n	'��		n	��		n	��		n	��	 	o	�� 			   g  nameg  
triplet-os� CYR+      h   �   ] [C  �       g  filenamef  system/base/target.scm�
	r
�� 		
  g  nameg  target-type�g  documentationf  <Return the GNU configuration triplet of the target platform.� CR7   h   �   ] 45 6     �       g  filenamef  system/base/target.scm�
	v
��		x	��		x	�� 		
  g  nameg  
target-cpu�g  documentationf  +Return the CPU name of the target platform.� CR\     h   �   ] 45 6     �       g  filenamef  system/base/target.scm�
	z
��		|	��		|	�� 		
  g  nameg  target-vendor�g  documentationf  .Return the vendor name of the target platform.� CRY       h   �   ] 45 6     �       g  filenamef  system/base/target.scm�
	~
��	 �	��	 �	�� 		
  g  nameg  	target-os�g  documentationf  8Return the operating system name of the target platform.� CR- h   �   ] [C  �       g  filenamef  system/base/target.scm�
 �
�� 		
  g  nameg  target-endianness�g  documentationf  4Return the endianness object of the target platform.� CR.      h   �   ] [C  �       g  filenamef  system/base/target.scm�
 �
�� 		
  g  nameg  target-word-size�g  documentationf  7Return the word size, in bytes, of the target platform.� CRC       g  m
		,  g  filenamef  system/base/target.scm�		
��	-	'	��	/	'	��	1	'	��	3	'	��	6	$
��	7	)	��	D	)
��	E	*	��	J	*	'��	R	*	��	U	*
��	V	+	��	c	+
���	-
��	4
��	<
��t	M
��	#	f
��
	i
��	m
���	r
���	v
���	z
��q	~
��< �
�� �
�� 	
   C6 