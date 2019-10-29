package colors

import (
  "github.com/nsf/termbox-go"
  "github.com/lucasb-eyer/go-colorful"
)

var hexPalette = map[termbox.Attribute]string{
  termbox.Attribute(0): "#000000",
  termbox.Attribute(1): "#000000",
  termbox.Attribute(2): "#800000",
  termbox.Attribute(3): "#008000",
  termbox.Attribute(4): "#808000",
  termbox.Attribute(5): "#000080",
  termbox.Attribute(6): "#800080",
  termbox.Attribute(7): "#008080",
  termbox.Attribute(8): "#c0c0c0",
  termbox.Attribute(9): "#808080",
  termbox.Attribute(10): "#ff0000",
  termbox.Attribute(11): "#00ff00",
  termbox.Attribute(12): "#ffff00",
  termbox.Attribute(13): "#0000ff",
  termbox.Attribute(14): "#ff00ff",
  termbox.Attribute(15): "#00ffff",
  termbox.Attribute(16): "#ffffff",
  termbox.Attribute(17): "#000000",
  termbox.Attribute(18): "#00005f",
  termbox.Attribute(19): "#000087",
  termbox.Attribute(20): "#0000af",
  termbox.Attribute(21): "#0000d7",
  termbox.Attribute(22): "#0000ff",
  termbox.Attribute(23): "#005f00",
  termbox.Attribute(24): "#005f5f",
  termbox.Attribute(25): "#005f87",
  termbox.Attribute(26): "#005faf",
  termbox.Attribute(27): "#005fd7",
  termbox.Attribute(28): "#005fff",
  termbox.Attribute(29): "#008700",
  termbox.Attribute(30): "#00875f",
  termbox.Attribute(31): "#008787",
  termbox.Attribute(32): "#0087af",
  termbox.Attribute(33): "#0087d7",
  termbox.Attribute(34): "#0087ff",
  termbox.Attribute(35): "#00af00",
  termbox.Attribute(36): "#00af5f",
  termbox.Attribute(37): "#00af87",
  termbox.Attribute(38): "#00afaf",
  termbox.Attribute(39): "#00afd7",
  termbox.Attribute(40): "#00afff",
  termbox.Attribute(41): "#00d700",
  termbox.Attribute(42): "#00d75f",
  termbox.Attribute(43): "#00d787",
  termbox.Attribute(44): "#00d7af",
  termbox.Attribute(45): "#00d7d7",
  termbox.Attribute(46): "#00d7ff",
  termbox.Attribute(47): "#00ff00",
  termbox.Attribute(48): "#00ff5f",
  termbox.Attribute(49): "#00ff87",
  termbox.Attribute(50): "#00ffaf",
  termbox.Attribute(51): "#00ffd7",
  termbox.Attribute(52): "#00ffff",
  termbox.Attribute(53): "#5f0000",
  termbox.Attribute(54): "#5f005f",
  termbox.Attribute(55): "#5f0087",
  termbox.Attribute(56): "#5f00af",
  termbox.Attribute(57): "#5f00d7",
  termbox.Attribute(58): "#5f00ff",
  termbox.Attribute(59): "#5f5f00",
  termbox.Attribute(60): "#5f5f5f",
  termbox.Attribute(61): "#5f5f87",
  termbox.Attribute(62): "#5f5faf",
  termbox.Attribute(63): "#5f5fd7",
  termbox.Attribute(64): "#5f5fff",
  termbox.Attribute(65): "#5f8700",
  termbox.Attribute(66): "#5f875f",
  termbox.Attribute(67): "#5f8787",
  termbox.Attribute(68): "#5f87af",
  termbox.Attribute(69): "#5f87d7",
  termbox.Attribute(70): "#5f87ff",
  termbox.Attribute(71): "#5faf00",
  termbox.Attribute(72): "#5faf5f",
  termbox.Attribute(73): "#5faf87",
  termbox.Attribute(74): "#5fafaf",
  termbox.Attribute(75): "#5fafd7",
  termbox.Attribute(76): "#5fafff",
  termbox.Attribute(77): "#5fd700",
  termbox.Attribute(78): "#5fd75f",
  termbox.Attribute(79): "#5fd787",
  termbox.Attribute(80): "#5fd7af",
  termbox.Attribute(81): "#5fd7d7",
  termbox.Attribute(82): "#5fd7ff",
  termbox.Attribute(83): "#5fff00",
  termbox.Attribute(84): "#5fff5f",
  termbox.Attribute(85): "#5fff87",
  termbox.Attribute(86): "#5fffaf",
  termbox.Attribute(87): "#5fffd7",
  termbox.Attribute(88): "#5fffff",
  termbox.Attribute(89): "#870000",
  termbox.Attribute(90): "#87005f",
  termbox.Attribute(91): "#870087",
  termbox.Attribute(92): "#8700af",
  termbox.Attribute(93): "#8700d7",
  termbox.Attribute(94): "#8700ff",
  termbox.Attribute(95): "#875f00",
  termbox.Attribute(96): "#875f5f",
  termbox.Attribute(97): "#875f87",
  termbox.Attribute(98): "#875faf",
  termbox.Attribute(99): "#875fd7",
  termbox.Attribute(100): "#875fff",
  termbox.Attribute(101): "#878700",
  termbox.Attribute(102): "#87875f",
  termbox.Attribute(103): "#878787",
  termbox.Attribute(104): "#8787af",
  termbox.Attribute(105): "#8787d7",
  termbox.Attribute(106): "#8787ff",
  termbox.Attribute(107): "#87af00",
  termbox.Attribute(108): "#87af5f",
  termbox.Attribute(109): "#87af87",
  termbox.Attribute(110): "#87afaf",
  termbox.Attribute(111): "#87afd7",
  termbox.Attribute(112): "#87afff",
  termbox.Attribute(113): "#87d700",
  termbox.Attribute(114): "#87d75f",
  termbox.Attribute(115): "#87d787",
  termbox.Attribute(116): "#87d7af",
  termbox.Attribute(117): "#87d7d7",
  termbox.Attribute(118): "#87d7ff",
  termbox.Attribute(119): "#87ff00",
  termbox.Attribute(120): "#87ff5f",
  termbox.Attribute(121): "#87ff87",
  termbox.Attribute(122): "#87ffaf",
  termbox.Attribute(123): "#87ffd7",
  termbox.Attribute(124): "#87ffff",
  termbox.Attribute(125): "#af0000",
  termbox.Attribute(126): "#af005f",
  termbox.Attribute(127): "#af0087",
  termbox.Attribute(128): "#af00af",
  termbox.Attribute(129): "#af00d7",
  termbox.Attribute(130): "#af00ff",
  termbox.Attribute(131): "#af5f00",
  termbox.Attribute(132): "#af5f5f",
  termbox.Attribute(133): "#af5f87",
  termbox.Attribute(134): "#af5faf",
  termbox.Attribute(135): "#af5fd7",
  termbox.Attribute(136): "#af5fff",
  termbox.Attribute(137): "#af8700",
  termbox.Attribute(138): "#af875f",
  termbox.Attribute(139): "#af8787",
  termbox.Attribute(140): "#af87af",
  termbox.Attribute(141): "#af87d7",
  termbox.Attribute(142): "#af87ff",
  termbox.Attribute(143): "#afaf00",
  termbox.Attribute(144): "#afaf5f",
  termbox.Attribute(145): "#afaf87",
  termbox.Attribute(146): "#afafaf",
  termbox.Attribute(147): "#afafd7",
  termbox.Attribute(148): "#afafff",
  termbox.Attribute(149): "#afd700",
  termbox.Attribute(150): "#afd75f",
  termbox.Attribute(151): "#afd787",
  termbox.Attribute(152): "#afd7af",
  termbox.Attribute(153): "#afd7d7",
  termbox.Attribute(154): "#afd7ff",
  termbox.Attribute(155): "#afff00",
  termbox.Attribute(156): "#afff5f",
  termbox.Attribute(157): "#afff87",
  termbox.Attribute(158): "#afffaf",
  termbox.Attribute(159): "#afffd7",
  termbox.Attribute(160): "#afffff",
  termbox.Attribute(161): "#d70000",
  termbox.Attribute(162): "#d7005f",
  termbox.Attribute(163): "#d70087",
  termbox.Attribute(164): "#d700af",
  termbox.Attribute(165): "#d700d7",
  termbox.Attribute(166): "#d700ff",
  termbox.Attribute(167): "#d75f00",
  termbox.Attribute(168): "#d75f5f",
  termbox.Attribute(169): "#d75f87",
  termbox.Attribute(170): "#d75faf",
  termbox.Attribute(171): "#d75fd7",
  termbox.Attribute(172): "#d75fff",
  termbox.Attribute(173): "#d78700",
  termbox.Attribute(174): "#d7875f",
  termbox.Attribute(175): "#d78787",
  termbox.Attribute(176): "#d787af",
  termbox.Attribute(177): "#d787d7",
  termbox.Attribute(178): "#d787ff",
  termbox.Attribute(179): "#d7af00",
  termbox.Attribute(180): "#d7af5f",
  termbox.Attribute(181): "#d7af87",
  termbox.Attribute(182): "#d7afaf",
  termbox.Attribute(183): "#d7afd7",
  termbox.Attribute(184): "#d7afff",
  termbox.Attribute(185): "#d7d700",
  termbox.Attribute(186): "#d7d75f",
  termbox.Attribute(187): "#d7d787",
  termbox.Attribute(188): "#d7d7af",
  termbox.Attribute(189): "#d7d7d7",
  termbox.Attribute(190): "#d7d7ff",
  termbox.Attribute(191): "#d7ff00",
  termbox.Attribute(192): "#d7ff5f",
  termbox.Attribute(193): "#d7ff87",
  termbox.Attribute(194): "#d7ffaf",
  termbox.Attribute(195): "#d7ffd7",
  termbox.Attribute(196): "#d7ffff",
  termbox.Attribute(197): "#ff0000",
  termbox.Attribute(198): "#ff005f",
  termbox.Attribute(199): "#ff0087",
  termbox.Attribute(200): "#ff00af",
  termbox.Attribute(201): "#ff00d7",
  termbox.Attribute(202): "#ff00ff",
  termbox.Attribute(203): "#ff5f00",
  termbox.Attribute(204): "#ff5f5f",
  termbox.Attribute(205): "#ff5f87",
  termbox.Attribute(206): "#ff5faf",
  termbox.Attribute(207): "#ff5fd7",
  termbox.Attribute(208): "#ff5fff",
  termbox.Attribute(209): "#ff8700",
  termbox.Attribute(210): "#ff875f",
  termbox.Attribute(211): "#ff8787",
  termbox.Attribute(212): "#ff87af",
  termbox.Attribute(213): "#ff87d7",
  termbox.Attribute(214): "#ff87ff",
  termbox.Attribute(215): "#ffaf00",
  termbox.Attribute(216): "#ffaf5f",
  termbox.Attribute(217): "#ffaf87",
  termbox.Attribute(218): "#ffafaf",
  termbox.Attribute(219): "#ffafd7",
  termbox.Attribute(220): "#ffafff",
  termbox.Attribute(221): "#ffd700",
  termbox.Attribute(222): "#ffd75f",
  termbox.Attribute(223): "#ffd787",
  termbox.Attribute(224): "#ffd7af",
  termbox.Attribute(225): "#ffd7d7",
  termbox.Attribute(226): "#ffd7ff",
  termbox.Attribute(227): "#ffff00",
  termbox.Attribute(228): "#ffff5f",
  termbox.Attribute(229): "#ffff87",
  termbox.Attribute(230): "#ffffaf",
  termbox.Attribute(231): "#ffffd7",
  termbox.Attribute(232): "#ffffff",
  termbox.Attribute(233): "#080808",
  termbox.Attribute(234): "#121212",
  termbox.Attribute(235): "#1c1c1c",
  termbox.Attribute(236): "#262626",
  termbox.Attribute(237): "#303030",
  termbox.Attribute(238): "#3a3a3a",
  termbox.Attribute(239): "#444444",
  termbox.Attribute(240): "#4e4e4e",
  termbox.Attribute(241): "#585858",
  termbox.Attribute(242): "#606060",
  termbox.Attribute(243): "#666666",
  termbox.Attribute(244): "#767676",
  termbox.Attribute(245): "#808080",
  termbox.Attribute(246): "#8a8a8a",
  termbox.Attribute(247): "#949494",
  termbox.Attribute(248): "#9e9e9e",
  termbox.Attribute(249): "#a8a8a8",
  termbox.Attribute(250): "#b2b2b2",
  termbox.Attribute(251): "#bcbcbc",
  termbox.Attribute(252): "#c6c6c6",
  termbox.Attribute(253): "#d0d0d0",
  termbox.Attribute(254): "#dadada",
  termbox.Attribute(255): "#e4e4e4",
  termbox.Attribute(256): "#eeeeee",
}

func MapTermboxColorToColor(color termbox.Attribute) (colorful.Color, error) {
  kolor, err := colorful.Hex(hexPalette[color])

  if err != nil {
    return colorful.Color{0,0,0}, err
  }

  return kolor, nil
}

func IsLightColor(color termbox.Attribute) bool {
  kolor, _ := colorful.Hex(hexPalette[color])
  _, _, brightness := kolor.Hsv()

  return brightness > 0.75
}
