GOOF----LE-4-2.0ú!      ] ^ 4        h      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  scripts¤	g  api-diff¤	 ¤		g  filenameS¤	
f  scripts/api-diff.scm¤	g  importsS¤	g  ice-9¤	g  common-list¤	 ¤	 ¤	g  format¤	 ¤	 ¤	g  getopt-long¤	 ¤	 ¤	 ¤	g  exportsS¤	 ¤	g  	autoloadsS¤	g  srfi¤	g  srfi-13¤	 ¤	g  string-tokenize¤	 ¤	 ¤	 g  set-current-module¤	!  ¤	"  ¤	#g  %include-in-guild-list¤	$f  ,Show differences between two scan-api files.¤	%g  %summary¤	&g  with-input-from-file¤	'g  read¤	(g  read-alist-file¤	)g  set-object-property!¤	*g  put¤	+g  object-property¤	,g  get¤	-g  assq-ref¤	.g  meta¤	/g  	interface¤	0g  groups¤	1g  make-hash-table¤	2g  for-each¤	3g  
hashq-set!¤	4g  read-api-alist-file¤	5g  	hashq-ref¤	6g  hang-by-the-roots¤	7g  set-difference¤	8g  diff?¤	9g  
diff+note!¤	:g  	hash-fold¤	;g  acons¤	<g  map¤	=g  car¤	>f  groups-removed: ~A
¤	?f  groups-added: ~A
¤	@g  length¤	Af   ~5@A  ~5@A  :  ¤	Bf  -¤	Cf  ~5@A ~5@A : ~5@A¤	Df  ~5@D ~5@D : ~5@D¤	Ef       ~A
¤	Fg  sort¤	Gg  union¤	Hg  string<?¤	Ig  symbol->string¤	Jg  details¤	Kf  ~A ~A:
¤	Lg  removals¤	Mf   ~A
¤	Ng  	additions¤	Of  ~A: no changes
¤	Pg  error¤	Qf  !api-diff: group-diff: bad options¤	Rg  
group-diff¤	Sg  single-char¤	TSd ¤	Ug  value¤	VU ¤	WJTV ¤	XW ¤	Yg  
option-ref¤	Zf  	/dev/null¤	[ZZ ¤	\g  string->symbol¤	]g  main¤C 5 h  À   ]4	
5	 4" >  "  G   #R$%R&'      h   O   ] 6   G       g  filenamef  scripts/api-diff.scm
	5			5	 		
   C      h   }   ] 6      u       g  file
		
  g  filenamef  scripts/api-diff.scm
	3
	
	4	 		
  g  nameg  read-alist-file C(R)i*R+i,R(-./*0123   h   k   ]L  6     c       g  group
		  g  filenamef  scripts/api-diff.scm
	@	'			A	>		A	) 		   C 
 h   |  ]!4 545454>  "  G  44	54	O 45>  "  G  >  "  G  C    t      g  file
		| g  alist			| g  meta			| g  	interface			| g  ht		E	l  g  filenamef  scripts/api-diff.scm
	:
		;				;			<			<			<			;			=			=	$		=			;		"	>		(	>		/	>		8	?		>	?		?	?	%	E	?		H	@		R	B	'	X	B	6	Z	B	'	_	@		q	?	 		|  g  nameg  read-api-alist-file C4R,02235       h   {   ]L L 4L 56 s       g  group
		  g  filenamef  scripts/api-diff.scm
	I			K	.		L	.		K	(		J	 		   C-0      h   w   ] L O 4 56o       g  x
		  g  filenamef  scripts/api-diff.scm
	H			M			M	&		M			I	 		   C      h0   ³   ]	4 54O  >  "  G   C    «       g  	interface
		, g  ht		)  g  filenamef  scripts/api-diff.scm
	F
		G				G			G			G			H	 		,  g  nameg  hang-by-the-roots C6R7    h   £   ]
4 5(  CC       g  a
		 g  b		 g  result			  g  filenamef  scripts/api-diff.scm
	Q
		R			R			S	 			  g  nameg  diff? C8R8    hp   S  ]H4 5$  4>  "  G  K"   4 5$  4>  "  G  K"   J$  6 C     K      g  a
		k g  b		k g  note-removals			k g  note-additions			k g  	note-same			k g  same?			k g  t			2 g  t		;	_  g  filenamef  scripts/api-diff.scm
	W
		X			Y			Y			Y	&	+	Y	8	3	Z		;	Z		D	Z	&	X	Z	9	e	[		i	[	 		k	  g  nameg  
diff+note! C9R6:;,0<=9> h   n   ] 6     f       g  removals
		  g  filenamef  scripts/api-diff.scm
	f			g	$		g	 		   C?  h   o   ] 6     g       g  	additions
		  g  filenamef  scripts/api-diff.scm
	h			i	$		i	 		   C       h   G   ] C    ?       g  filenamef  scripts/api-diff.scm
	j	 		
   C2-@ABC9@  h   j   ]4 5N C   b       g  subs
		  g  filenamef  scripts/api-diff.scm
	x	#		y	5		y	% 		   C@        h   j   ]4 5N C   b       g  adds
		  g  filenamef  scripts/api-diff.scm
	z	#		{	5		{	% 		   C    h   G   ] C    ?       g  filenamef  scripts/api-diff.scm
	|	# 		
   CDE  h   /  ]94L  54L 5$  45"  $  45"  $  $  	"  "  4$  "  $  "  >  "  G  "  4>  "  nG  "  g$  ]$  O

HH4O 	O 
>  "  G  4J
J>  "  G  "  "ÿÿ"  "ÿÿ| 6   '      g  group
	 g  old	 g  new		 g  	old-count		* g  	new-count		> g  delta		[ g  	add-count	 » g  	sub-count	 »  g  filenamef  scripts/api-diff.scm
	k			l	#		l			m	#		l			n	)		n	2	*	l		2	o	)	3	o	2	>	l		F	p	%	Q	p	2	[	l		^	q		c	q	$	i	r	!	q	r	/	w	s	!		s	/ 	q	  	   	+  	>  	B  	F ¡ 	  ­	t	 ¹	t	  »	u	  À	v	" ä	}	" é	}	- ì	~	4 ö	}	" 	$ 	 &	   CFGHI     h      ]4 5456    z       g  a
		 g  b		  g  filenamef  scripts/api-diff.scm
 		 	'	 	'	 	 			   C-J-9KL2M    h   j   ] 6     b       g  x
		  g  filenamef  scripts/api-diff.scm
 	4	 	A	 	6 		   C   h(      ]4L >  "  G   6              g  removals
		!  g  filenamef  scripts/api-diff.scm
 	*	 	*	 	5	 	1	 	*	! 	* 		!   CKN2M   h   j   ] 6     b       g  x
		  g  filenamef  scripts/api-diff.scm
 	4	 	A	 	6 		   C   h(      ]4L >  "  G   6              g  	additions
		!  g  filenamef  scripts/api-diff.scm
 	*	 	*	 	5	 	1	 	*	! 	* 		!   CO   h   Z   ] L 6     R       g  filenamef  scripts/api-diff.scm
 	*	 	7	 	, 		
   C   h`   ë   ]4L  5$  "  4L 5$  "   O  O  O 6       ã       g  group
		Y g  t		 g  old		Y g  t		(	: g  new		:	Y  g  filenamef  scripts/api-diff.scm
 		 	,	 	(	 	C	 		  	,	( 	(	7 	C	: 		Y 	 		Y   CPQ      hÀ   a  - 1 3 
4 544554545445545(  74	
>  "  G  O 4455645		$  O 	66      Y      g  i-old
		 º g  i-new		 º g  options			 º g  i-old		 º g  g-old		# º g  g-old-names		. º g  i-new		7 º g  g-new		J º g  g-new-names		U º g  t		  º  
g  filenamef  scripts/api-diff.scm
	]
	
	^			^			_			_	!		_	%		_	0	!	_	%	#	_		#	^		&	`		.	^		1	a		7	^		:	b		?	b	!	@	b	%	F	b	0	H	b	%	J	b		J	^		M	c		U	^		]	d		^	e	  	  	  	 	k	  	  	  	 	d	 ´ 	 ¸ 	 º 	 %		 º	
	  g  nameg  
group-diff CRRXY[4J<\R  hx     -  1  3 4 5454545H45$  4	4
,55JK"   J@  ü      g  args
			v g  p		v g  rest		!	v g  i-old		+	v g  i-new		6	v g  options		9	v g  t		E	l  g  filenamef  scripts/api-diff.scm
 
	
 		 		 		 		 		 		 ¢		 ¢		 ¢	!	! ¢		! 		$ £		) £	%	+ £		+ 		. ¤		3 ¤	%	6 ¤		6 		9 ¥		9 		< ¦		B ¦		E ¦		E ¦		O ¨	*	P ©	*	U ª	/	_ ©	*	` ¨	$	c ¨		e ¨		v ®	 #			v


  g  nameg  api-diff CRi]RC¸       g  m
		0  g  filenamef  scripts/api-diff.scm		)
	4	0
	6	1		9	1
H	3
O	7
V	8
ÿ	:
N	F
	Q
ö	W
ì	]
 
 °
 	
   C6 